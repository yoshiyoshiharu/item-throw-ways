package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
  var err error

  conn := "academy02:Zo.EbZjFLQxaNqs3o3M*@tcp(ca-go-academy-3.c9ml7do7yvmn.ap-northeast-1.rds.amazonaws.com)/academy02?parseTime=true"
  Db, err = sql.Open("mysql", conn)
  if err != nil {
    log.Println("DB接続失敗")
    log.Fatal(err)
  }

  err = Db.Ping()
  if err != nil {
    log.Println("DBPing失敗")
    log.Fatal(err)
  }
  log.Println("DB接続成功")
}

