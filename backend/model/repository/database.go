package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
  var err error

  conn := "user:password@tcp(db:3306)/db?charset=utf8&parseTime=True"
  Db, err = sql.Open("mysql", conn)
  if err != nil {
    log.Fatal(err)
  }

  err = Db.Ping()
  if err != nil {
    log.Fatal(err)
  }
  log.Println("DB接続成功")
}

