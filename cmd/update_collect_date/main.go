package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
	"github.com/yoshiyoshiharu/item-throw-ways/pkg/date"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/05syusyubi/syusyubi.csv"
)

var kindRepository = repository.NewKindRepository()
var itemRepository = repository.NewItemRepository()
var kinds = kindRepository.GetKinds()

func main() {
	updateCollectDateFromCsv()
}

type CollectDate struct {
  Weekday time.Weekday
  n int
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
    // ヘッダー行はスキップ
		if i == 0 {
			continue
		}

    area_id := i + 1
    town := row[0]
    street := row[1]
    kanen := row[2]
    funen := row[2]
    shigen := row[3]
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

func splitNthWeekday (string) {
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

