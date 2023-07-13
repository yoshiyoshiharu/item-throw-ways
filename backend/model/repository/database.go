package repository

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

var (
  DB_USER = os.Getenv("DB_USER")
  DB_PASSWORD = os.Getenv("DB_PASSWORD")
  DB_HOST = os.Getenv("DB_HOST")
  DB_NAME = os.Getenv("DB_NAME")
)

func init() {
  var err error

  dsn := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
  fmt.Println(dsn)
  Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
   if err != nil {
    log.Println("DB接続失敗")
    log.Fatal(err)
  }

  log.Println("DB接続成功")
}

