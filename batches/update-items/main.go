package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/06bunbetuhinmoku/bunbetuhinmoku.csv"
)

var Kinds []entity.Kind
var kindRepository = repository.NewKindRepository()
var itemRepository = repository.NewItemRepository()

func main() {
	kinds := kindRepository.GetKinds()
	updateItemsFromCsv(kinds)
}

func updateItemsFromCsv(kinds []entity.Kind) {
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

	tx, err := repository.Db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = repository.Db.Query("DELETE FROM items;")
	_, err = repository.Db.Query("DELETE FROM item_kinds;")
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
		if i == 0 || itemRepository.ItemExists(item_name) {
			continue
		}

		fmt.Println(item_id, item_name, kind_names, price, remarks)

		_, err = repository.Db.Exec("INSERT INTO items (id, name, price, remarks) VALUES (?, ?, ?, ?)", item_id, item_name, price, remarks)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		for _, kind_name := range kind_names {
			kind_id, err := kindRepository.GetKindIdByName(kinds, kind_name)
			if err != nil {
				tx.Rollback()
				log.Fatal(err)
			}
			_, err = repository.Db.Exec("INSERT INTO item_kinds (item_id, kind_id) VALUES (?, ?)", item_id, kind_id)
		}
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func GetKindsFromCell(str string) []string {
	return strings.Split(str, "、")
}

