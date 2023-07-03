package main

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/05syusyubi/syusyubi.csv"
)

var Db *sql.DB
var Kinds []Kind

type Kind struct {
  Id int
  Name string
}

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
  SetKinds()
	updateItemsFromCsv()
}

func updateItemsFromCsv() {
	resp, err := http.Get(
		API_URL,
	)
	if err != nil {
    log.Fatal(err)
	}
	defer resp.Body.Close()

	r := csv.NewReader(transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder()))
	rows, err := r.ReadAll()
	if err != nil {
    log.Fatal(err)
	}

  tx, err := Db.Begin()
  if err != nil {
    log.Fatal(err)
  }

  _, err = Db.Query("DELETE FROM items;")
  _, err = Db.Query("DELETE FROM item_kinds;")
  if err != nil {
    tx.Rollback()
    log.Fatal(err)
  }

	for i, row := range rows {
    item_id := i + 1
    item_name := row[1]
    kind_names := GetKindsFromCell(row[2])
    price, _ := strconv.Atoi(row[3])
    remarks := row[4]

    // ヘッダー行はスキップ
    if i == 0 || ItemExists(Db, item_name){
      continue
    }

    fmt.Println(item_id, item_name, kind_names, price, remarks)

    _, err = Db.Exec("INSERT INTO items (id, name, price, remarks) VALUES (?, ?, ?, ?)", item_id, item_name, price, remarks)
    if err != nil {
      tx.Rollback()
      log.Fatal(err)
    }

    for _, kind_name := range kind_names {
      kind_id, err := GetKindIdByName(kind_name)
      if err != nil {
        tx.Rollback()
        log.Fatal(err)
      }
      _, err = Db.Exec("INSERT INTO item_kinds (item_id, kind_id) VALUES (?, ?)", item_id, kind_id)
    }
	}

  if err = tx.Commit(); err != nil {
    log.Fatal(err)
  }
}

func SetKinds() {
  rows, err := Db.Query("SELECT id, name FROM kinds;")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    var kind Kind
    err := rows.Scan(&kind.Id, &kind.Name)
    if err != nil {
      log.Fatal(err)
    }
    Kinds = append(Kinds, kind)
  }
}

func GetKindIdByName(name string) (int, error) {
  for _, kind := range Kinds {
    if kind.Name == name {
      return kind.Id, nil
    }
  }

  return 0, errors.New("Not found")
}

func GetKindsFromCell(str string) []string {
  return strings.Split(str, "、")
}

func ItemExists(db * sql.DB, name string) bool {
    sqlStmt := `SELECT name FROM items WHERE name = ?`
    err := db.QueryRow(sqlStmt, name).Scan(&name)
    if err != nil {
        if err != sql.ErrNoRows {
            log.Fatal(err)
        }
        return false
    }
    return true
}
