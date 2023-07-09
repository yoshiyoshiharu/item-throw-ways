package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/06bunbetuhinmoku/bunbetuhinmoku.csv"
)

var itemRepository = repository.NewItemRepository()

func handler(c context.Context) {
  updateItemsFromCsv()
}

func main() {
  lambda.Start(handler)
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

	repository.Db.Exec("DELETE FROM items;")
	repository.Db.Exec("DELETE FROM item_kinds;")

	for i, row := range rows {
		item_id := i
		item_name := row[1]
		kind_names := GetKindsFromCell(row[2])
		price, _ := strconv.Atoi(row[3])
		remarks := row[4]

    fmt.Println(item_id, item_name, kind_names, price, remarks)

		// ヘッダー行はスキップ
		if i == 0 || itemRepository.ItemExists(item_name) {
			continue
		}

    item := entity.Item{Id: item_id, Name: item_name, Price: price, Remarks: remarks}
    var kinds []entity.Kind

    repository.Db.Find(&kinds, "name IN ?", kind_names)

    item.Kinds = kinds

    repository.Db.Create(&item)
	}
}

func GetKindsFromCell(str string) []string {
	return strings.Split(str, "、")
}

