package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/pkg/database"
	"gorm.io/gorm"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/06bunbetuhinmoku/bunbetuhinmoku.csv"
  HIRAGANA_TRANSLATION_API_URL = "https://labs.goo.ne.jp/api/hiragana"
  CONCURRENCY = 10
)

var (
  items []entity.Item
  allKinds []entity.Kind
  mu sync.Mutex
  wg sync.WaitGroup
)

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

func init() {
  database.Db.Find(&allKinds)
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

  itemChan := make(chan *entity.Item)
  semaphore := make(chan struct{}, CONCURRENCY)

  insertId := 0
	for i, row := range rows {
    // ヘッダー行はスキップ
    if i == 0 {
      continue
    }

    wg.Add(1)
    semaphore <- struct{}{} // セマフォに空きがでるまでブロック

    insertId++
    go func(insertId int, row []string) {
      itemId := insertId
      itemName := row[1]
      kindNames := GetKindsFromCell(row[2])
      price, _ := strconv.Atoi(row[3])
      remarks := row[4]

      itemNameKana, err := TranslateToHiragana(itemName)
      if err != nil {
        log.Fatal(err)
      }

      var kinds []entity.Kind
      for _, kindName := range kindNames {
        kind := findKind(kindName, allKinds)

        kinds = append(kinds, kind)
      }

      itemChan <- entity.NewItem(
        itemId,
        itemName,
        itemNameKana,
        price,
        remarks,
        kinds,
      )

      <- semaphore
      wg.Done()
    }(insertId, row)

    go func() {
      for item := range itemChan {
        mu.Lock()
        if itemExists(item.Name, items) {
          mu.Unlock()
          continue
        }
        items = append(items, *item)
        mu.Unlock()
      }
      close(itemChan)
    }()
  }

  wg.Wait()

  database.Db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Exec("DELETE FROM items").Error; err != nil {
      return err
    }
    if err := tx.Exec("DELETE FROM item_kinds").Error; err != nil {
      return err
    }
    if err := tx.Create(&items).Error; err != nil {
      return err
    }
    return nil
  })
}

func GetKindsFromCell(str string) []string {
	return strings.Split(str, "、")
}

func TranslateToHiragana(name string) (string, error) {
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

  return responseBody.Converted, nil
}

func itemExists(name string, items []entity.Item) bool {
  for _, item := range items {
    if item.Name == name {
      return true
    }
  }
  return false
}

func findKind(kindName string, allKinds []entity.Kind) entity.Kind {
  for _, kind := range allKinds {
    if kind.Name == kindName {
      return kind
    }
  }

  return entity.Kind{}
}

