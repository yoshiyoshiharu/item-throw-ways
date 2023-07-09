package repository

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
  var err error

  dsn := "academy02:Zo.EbZjFLQxaNqs3o3M*@tcp(ca-go-academy-3.c9ml7do7yvmn.ap-northeast-1.rds.amazonaws.com)/academy02?parseTime=true"
  Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
   if err != nil {
    log.Println("DB接続失敗")
    log.Fatal(err)
  }

  log.Println("DB接続成功")
}

