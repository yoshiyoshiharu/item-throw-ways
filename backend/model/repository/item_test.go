package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

func TestItemRepository_FindAll(t *testing.T) {
  gormDB, mock := newMockDB()

  repo := NewItemRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Item 1").
		AddRow(2, "Item 2")

	mock.ExpectQuery("SELECT \\* FROM `items`").WillReturnRows(rows)

	items := repo.FindAll()

  t.Run("[正常系] FindAllは全件返すこと", func(t *testing.T) {
    assert.Equal(t, 2, len(items))
    assert.Equal(t, "Item 1", items[0].Name)
    assert.Equal(t, "Item 2", items[1].Name)

    assert.NoError(t, mock.ExpectationsWereMet())
  })
}

func TestItemRepository_DeleteAndAll(t *testing.T) {
  gormDB, mock := newMockDB()

  repo := NewItemRepository(gormDB)

  t.Run("[正常系] itemsとitem_kindsを全消去し、itemsを挿入すること", func(t *testing.T) {
    item1 := &entity.Item{ID: 1, Name: "Item 1"}
    item2 := &entity.Item{ID: 2, Name: "Item 2"}

    items := []*entity.Item{
      item1,
      item2,
    }

    mock.ExpectBegin()
    mock.ExpectExec("DELETE FROM items").WillReturnResult(sqlmock.NewResult(0, 0))
    mock.ExpectExec("DELETE FROM item_kinds").WillReturnResult(sqlmock.NewResult(0, 0))
    mock.ExpectExec(
      regexp.QuoteMeta("INSERT INTO `items` (`name`,`name_kana`,`price`,`remarks`,`id`) VALUES (?,?,?,?,?),(?,?,?,?,?)")).
      WithArgs(item1.Name, item1.NameKana, item1.Price, item1.Remarks, item1.ID, item2.Name, item2.NameKana, item2.Price, item2.Remarks, item2.ID).
      WillReturnResult(sqlmock.NewResult(0, 2))
    mock.ExpectCommit()

    err := repo.DeleteAndInsertAll(items)

    assert.NoError(t, err)
    assert.NoError(t, mock.ExpectationsWereMet())
  })

  t.Run("[異常系] トランザクションでロールバックした場合は、エラーを返すこと", func(t *testing.T) {
    item1 := &entity.Item{ID: 1, Name: "Item 1"}
    invalidItem := &entity.Item{ID: 1, Name: "Item 2"}

    items := []*entity.Item{
      item1,
      invalidItem,
    }

    mock.ExpectBegin()
    mock.ExpectExec("DELETE FROM items").WillReturnResult(sqlmock.NewResult(0, 0))
    mock.ExpectExec("DELETE FROM item_kinds").WillReturnResult(sqlmock.NewResult(0, 0))
    mock.ExpectExec(
      regexp.QuoteMeta("INSERT INTO `items` (`name`,`name_kana`,`price`,`remarks`,`id`) VALUES (?,?,?,?,?),(?,?,?,?,?)")).
      WithArgs(item1.Name, item1.NameKana, item1.Price, item1.Remarks, item1.ID, invalidItem.Name, invalidItem.NameKana, invalidItem.Price, invalidItem.Remarks, invalidItem.ID).
      WillReturnResult(sqlmock.NewResult(0, 2))
    mock.ExpectRollback()

    err := repo.DeleteAndInsertAll(items)

    assert.Error(t, err)
    assert.Error(t, mock.ExpectationsWereMet())
  })
}
