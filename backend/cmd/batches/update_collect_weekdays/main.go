package main

import (
	"context"
	"encoding/csv"
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
	updateCollectWeekdayFromCsv()
}

func main() {
	lambda.Start(handler)
}

type CollectWeekday struct {
	Weekday time.Weekday
	Lap     int
}

func updateCollectWeekdayFromCsv() {
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
	repository.Db.Exec("DELETE FROM area_collect_weekdays;")

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
		shigenWeekday, err := date.JaWeekdayToEn(shigen)
		if err != nil {
			log.Fatal(err)
		}

		kanenKind := entity.Kind{}
		repository.Db.Where("name = ?", "可燃ごみ").Limit(1).Find(&kanenKind)
		kindWeekdays := map[string][]CollectWeekday{
			"可燃ごみ": kanenWeekdays,
			"不燃ごみ": funenWeekdays,
			"資源":   []CollectWeekday{{shigenWeekday, 0}},
		}
		for kindName, weekdays := range kindWeekdays {
			kind := entity.Kind{}
			repository.Db.Where("name = ?", kindName).Limit(1).Find(&kind)
			for _, weekday := range weekdays {
				areaCollectWeekday := entity.AreaCollectWeekday{Area: area, Kind: kind, Weekday: int(weekday.Weekday), Lap: weekday.Lap}
				repository.Db.Create(&areaCollectWeekday)
			}
		}
	}
}

func splitWeekday(str string) []CollectWeekday {
	var collectWeekdays []CollectWeekday
	weekdays := strings.Split(str, "・")

	for _, weekday := range weekdays {
		weekday, err := date.JaWeekdayToEn(weekday)
		if err != nil {
			log.Fatal(err)
		}

		collectWeekdays = append(collectWeekdays, CollectWeekday{weekday, 0})
	}

	return collectWeekdays
}

func splitNthWeekday(str string) []CollectWeekday {
	var collectWeekdays []CollectWeekday
	// 第2・第4火曜日 から nとweekdayを抜き出す
	reg := regexp.MustCompile(`第(\d+)・第(\d+)(\D+)`)
	matches := reg.FindStringSubmatch(str)

	n1, err := strconv.Atoi(matches[1])
	n2, err := strconv.Atoi(matches[2])
	weekday, err := date.JaWeekdayToEn(matches[3])
	if err != nil {
		log.Fatal(err)
	}

	collectWeekdays = append(collectWeekdays, CollectWeekday{weekday, n1})
	collectWeekdays = append(collectWeekdays, CollectWeekday{weekday, n2})

	return collectWeekdays
}
