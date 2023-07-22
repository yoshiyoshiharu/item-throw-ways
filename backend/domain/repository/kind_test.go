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
		AddRow(1, "Kind 1").
		AddRow(2, "Kind 2")

	mock.ExpectQuery("SELECT \\* FROM `kinds`").WillReturnRows(rows)

	kinds := repo.FindAll()

	t.Run("[正常系] FindAllは全件返すこと", func(t *testing.T) {
		assert.Equal(t, 2, len(kinds))
		assert.Equal(t, "Kind 1", kinds[0].Name)
		assert.Equal(t, "Kind 2", kinds[1].Name)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
