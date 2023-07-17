package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestKindRepository_FindAll(t *testing.T) {
  gormDB, mock := newMockDB()

  repo := NewKindRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "可燃ごみ").
		AddRow(2, "不燃ごみ")

	mock.ExpectQuery("SELECT \\* FROM `kinds`").WillReturnRows(rows)

	areas := repo.FindAll()

  t.Run("[正常系] FindAllは全件返すこと", func(t *testing.T) {
    assert.Equal(t, 2, len(areas))
    assert.Equal(t, "可燃ごみ", areas[0].Name)
    assert.Equal(t, "不燃ごみ", areas[1].Name)

    assert.NoError(t, mock.ExpectationsWereMet())
  })
}

