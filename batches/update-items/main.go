package main

import (
	"database/sql"
	"encoding/csv"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/06bunbetuhinmoku/bunbetuhinmoku.csv"
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

func main() {
	updateItemsFromCsv()
}

func updateItemsFromCsv() error {
	resp, err := http.Get(
		API_URL,
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r := csv.NewReader(transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder()))
	rows, err := r.ReadAll()
	if err != nil {
		return err
	}

  tx, err := Db.Begin()
  if err != nil {
    return err
  }

  _, err = Db.Query("DELETE FROM items;")
  if err != nil {
    tx.Rollback()
    return err
  }

	for i, row := range rows {
    id := i + 1
    item := row[1]
    // kind := row[2]
    // price := row[3]
    // remarks := row[4]

    _, err = Db.Exec("INSERT INTO items (id, name) VALUES (?, ?)", id, item)
    if err != nil {
      tx.Rollback()
      return err
    }
	}

  if err := tx.Commit(); err != nil {
    log.Fatal(err)
  }

	return nil
}

