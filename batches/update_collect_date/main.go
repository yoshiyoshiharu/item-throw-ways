package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"strconv"

	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/05syusyubi/syusyubi.csv"
)

var Kinds []entity.Kind
var kindRepository = repository.NewKindRepository()
var itemRepository = repository.NewItemRepository()

func main() {
	kinds := kindRepository.GetKinds()
	updateCollectDateFromCsv(kinds)
}

func updateCollectDateFromCsv(kinds []entity.Kind) {
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

	_, err = repository.Db.Query("DELETE FROM areas;")
	_, err = repository.Db.Query("DELETE FROM area_collect_dates;")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	for i, row := range rows {
		item_id := i + 1
		item_name := row[1]
		price, _ := strconv.Atoi(row[3])
		remarks := row[4]

		// ヘッダー行はスキップ
		if i == 0 {
			continue
		}

		_, err = repository.Db.Exec("INSERT INTO items (id, name, price, remarks) VALUES (?, ?, ?, ?)", item_id, item_name, price, remarks)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

