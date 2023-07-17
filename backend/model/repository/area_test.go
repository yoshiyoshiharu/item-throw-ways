package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAreaRepository_FindAll(t *testing.T) {
  gormDB, mock := newMockDB()

  repo := NewAreaRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Area 1").
		AddRow(2, "Area 2")

	mock.ExpectQuery("SELECT \\* FROM `areas`").WillReturnRows(rows)

	areas := repo.FindAll()

  t.Run("[正常系] FindAllは全件返すこと", func(t *testing.T) {
    assert.Equal(t, 2, len(areas))
    assert.Equal(t, "Area 1", areas[0].Name)
    assert.Equal(t, "Area 2", areas[1].Name)

    assert.NoError(t, mock.ExpectationsWereMet())
  })
}


func TestAreaRepository_FindById(t *testing.T) {
  gormDB, mock := newMockDB()

  repo := NewAreaRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Area 1")

	mock.ExpectQuery("SELECT \\* FROM `areas` WHERE id = ?").WithArgs(1).WillReturnRows(rows)

	area, _ := repo.FindById(1)

  t.Run("[正常系] FindByIdは指定したIDのエリアを返すこと", func(t *testing.T) {
    assert.Equal(t, "Area 1", area.Name)

    assert.NoError(t, mock.ExpectationsWereMet())
  })
}
