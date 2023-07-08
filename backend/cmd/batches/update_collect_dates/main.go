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
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
	date "github.com/yoshiyoshiharu/item-throw-ways/pkg"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/05syusyubi/syusyubi.csv"
)

var kindRepository = repository.NewKindRepository()
var itemRepository = repository.NewItemRepository()
var kinds = kindRepository.GetKinds()

func handler(c context.Context) {
  updateCollectDateFromCsv()
}

func main() {
  lambda.Start(handler)
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

    area_id := i
    town := row[0]
    street := row[1]
    kanen := row[2]
    funen := row[3]
    shigen := row[4]

    repository.Db.Exec("INSERT INTO areas (id, name) VALUES (?, ?)", area_id, town + street)
    kanenDates := splitWeekday(kanen)
    funenDates := splitNthWeekday(funen)
    jaShigen, err := date.JaWeekdayToEn(shigen); if err != nil {
      log.Fatal(err)
    }

    shigenDate := CollectDate{jaShigen, 0}

    currentDate := time.Now()
    for _, kanenDate := range kanenDates {
      allWeekdayDate := date.AllWeekdayDates(kanenDate.Weekday, currentDate.Year(), currentDate.Month()); if err != nil {
        log.Fatal(err)
      }
      insertCollectDate(area_id, "可燃ごみ", allWeekdayDate)
    }

    for _, funenDate := range funenDates {
      nthWeekdayDate, err := date.NthWeekdayDate(funenDate.n, funenDate.Weekday, currentDate.Year(), currentDate.Month()); if err != nil {
        log.Fatal(err)
      }
      insertCollectDate(area_id, "不燃ごみ", []time.Time{nthWeekdayDate})
    }

    allWeekdayDate := date.AllWeekdayDates(shigenDate.Weekday, currentDate.Year(), currentDate.Month()); if err != nil {
      log.Fatal(err)
    }
    insertCollectDate(area_id, "資源", allWeekdayDate)
	}
}

func insertCollectDate(area_id int, kindName string, dates []time.Time) {
  kindId, err := kindRepository.GetKindIdByName(kinds, kindName); if err != nil {
    log.Fatal(err)
  }
  stmt, _ := repository.Db.Prepare("INSERT INTO area_collect_dates (area_id, kind_id, date) VALUES (?, ?, ?)")

  for _, date := range dates {
    fmt.Println("insertCollectDate", area_id, kindId, date)
    stmt.Exec(area_id, kindId, date)
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

