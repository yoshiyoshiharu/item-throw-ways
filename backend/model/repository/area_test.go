package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func TestFindAll(t *testing.T) {
  db, mock, _ := sqlmock.New()
  defer db.Close()

  gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
    SkipInitializeWithVersion: true,
	}), &gorm.Config{})

  repo := NewAreaRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "テストエリア1").
		AddRow(2, "テストエリア2")

	mock.ExpectQuery("SELECT \\* FROM `areas`").WillReturnRows(rows)

	areas := repo.FindAll()

	assert.Equal(t, 2, len(areas))
	assert.Equal(t, "テストエリア1", areas[0].Name)
	assert.Equal(t, "テストエリア2", areas[1].Name)

	assert.NoError(t, mock.ExpectationsWereMet())
}
