package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
	date "github.com/yoshiyoshiharu/item-throw-ways/pkg"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/05syusyubi/syusyubi.csv"
)

var itemRepository = repository.NewItemRepository()

func handler(c context.Context) {
  updateCollectDateFromCsv()
}

func main() {
  lambda.Start(handler)
}

type CollectDate struct {
  Weekday time.Weekday
  N int
}

func updateCollectDateFromCsv() {
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

	repository.Db.Exec("DELETE FROM areas;")
	repository.Db.Exec("DELETE FROM area_collect_dates;")

	for i, row := range rows {
    // ヘッダー行はスキップ
		if i == 0 {
			continue
		}

    area_id := i
    town := row[0]
    street := row[1]
    kanen := row[2]
    funen := row[3]
    shigen := row[4]

    area := entity.Area{Id: area_id, Name: town + street}
    repository.Db.Create(&area)

    kanenWeekdays := splitWeekday(kanen)
    funenWeekdays := splitNthWeekday(funen)
    shigenWeekday, err := date.JaWeekdayToEn(shigen); if err != nil {
      log.Fatal(err)
    }

    fmt.Println("可燃ごみ", kanenWeekdays)
    fmt.Println("不燃ごみ", funenWeekdays)
    fmt.Println("資源ごみ", shigenWeekday)
  }
}

func splitWeekday (str string) []CollectDate {
  var collectDates []CollectDate
  weekdays := strings.Split(str, "・")

  for _, weekday := range weekdays {
    weekday, err := date.JaWeekdayToEn(weekday)
    if err != nil {
      log.Fatal(err)
    }

    collectDates = append(collectDates, CollectDate{weekday, 0})
  }

  return collectDates
}

func splitNthWeekday (str string) []CollectDate {
  var collectDates []CollectDate
  // 第2・第4火曜日 から nとweekdayを抜き出す
  reg := regexp.MustCompile(`第(\d+)・第(\d+)(\D+)`)
  matches := reg.FindStringSubmatch(str)

  n1, err := strconv.Atoi(matches[1])
  n2, err := strconv.Atoi(matches[2])
  weekday, err := date.JaWeekdayToEn(matches[3])
  if err != nil {
    log.Fatal(err)
  }

  collectDates = append(collectDates, CollectDate{weekday, n1})
  collectDates = append(collectDates, CollectDate{weekday, n2})

  return collectDates
}

