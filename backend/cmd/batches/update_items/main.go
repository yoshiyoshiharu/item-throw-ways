package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/06bunbetuhinmoku/bunbetuhinmoku.csv"
  HIRAGANA_TRANSLATION_API_URL = "https://labs.goo.ne.jp/api/hiragana"
)

var itemRepository = repository.NewItemRepository()

type RequestBody struct {
  AppId string `json:"app_id"`
  OutputType string `json:"output_type"`
  Sentence string `json:"sentence"`
}

type ResponseBody struct {
  Converted string `json:"converted"`
}

func handler(c context.Context) {
  updateItemsFromCsv()
}

func main() {
  lambda.Start(handler)
}

func updateItemsFromCsv() {
  startBatch := time.Now()
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
    item_kana, err := TranslateToHiragana(item_name)
		kind_names := GetKindsFromCell(row[2])
		price, _ := strconv.Atoi(row[3])
		remarks := row[4]
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println(item_id, item_name, item_kana, kind_names, price, remarks)

		// ヘッダー行はスキップ
    start := time.Now()
		if i == 0 || itemRepository.ItemExists(item_name) {
			continue
		}
    fmt.Println("ItemExist: ", time.Now().Sub(start))

    item := entity.Item{ID: item_id, Name: item_name, NameKana: item_kana, Price: price, Remarks: remarks}
    var kinds []entity.Kind

    start = time.Now()
    repository.Db.Find(&kinds, "name IN ?", kind_names)
    fmt.Println("FindKinds: ", time.Now().Sub(start))

    item.Kinds = kinds

<<<<<<< Updated upstream
    start = time.Now()
    repository.Db.Create(&item)
    fmt.Println("CreateItem: ", time.Now().Sub(start))
	}
=======
    items = append(items, item)
  }

  repository.Db.Create(&items)
  fmt.Println("Batch: ", time.Now().Sub(startBatch))
>>>>>>> Stashed changes
}

func GetKindsFromCell(str string) []string {
	return strings.Split(str, "、")
}

func TranslateToHiragana(name string) (string, error) {
  start := time.Now()

  requestBody := &RequestBody{
    AppId: os.Getenv("HIRAGANA_TRANSLATION_APP_ID"),
    OutputType: "hiragana",
    Sentence: name,
  }

  jsonString, err := json.Marshal(requestBody)
  if err != nil {
    return "", err
  }

  req, _ := http.NewRequest("POST", HIRAGANA_TRANSLATION_API_URL, bytes.NewBuffer(jsonString))
  req.Header.Set("Content-Type", "application/json")
  client := new(http.Client)
  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }

  defer resp.Body.Close()

  var responseBody ResponseBody
  json.NewDecoder(resp.Body).Decode(&responseBody)

  fmt.Println("TranslateToHiragana: ", time.Now().Sub(start))

  return responseBody.Converted, nil
}
