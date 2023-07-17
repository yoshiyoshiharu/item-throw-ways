package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newMockDB() (*gorm.DB, sqlmock.Sqlmock) {
  db, mock, _ := sqlmock.New()

  gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
    SkipInitializeWithVersion: true,
	}), &gorm.Config{})
  
  return gormDB, mock
}

func TestFindAll(t *testing.T) {
  gormDB, mock := newMockDB()

  repo := NewAreaRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "テストエリア1").
		AddRow(2, "テストエリア2")

	mock.ExpectQuery("SELECT \\* FROM `areas`").WillReturnRows(rows)

	areas := repo.FindAll()

  t.Run("[正常系] FindAllは全件返すこと", func(t *testing.T) {
    assert.Equal(t, 2, len(areas))
    assert.Equal(t, "テストエリア1", areas[0].Name)
    assert.Equal(t, "テストエリア2", areas[1].Name)

    assert.NoError(t, mock.ExpectationsWereMet())
  })
}


func TestFindById(t *testing.T) {
  gormDB, mock := newMockDB()

  repo := NewAreaRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "テストエリア1")

	mock.ExpectQuery("SELECT \\* FROM `areas` WHERE id = ?").WithArgs(1).WillReturnRows(rows)

	area, _ := repo.FindById(1)

  t.Run("[正常系] FindByIdは指定したIDのエリアを返すこと", func(t *testing.T) {
    assert.Equal(t, "テストエリア1", area.Name)

    assert.NoError(t, mock.ExpectationsWereMet())
  })
}
