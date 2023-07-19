package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
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


