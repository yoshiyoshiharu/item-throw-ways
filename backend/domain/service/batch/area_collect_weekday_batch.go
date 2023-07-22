package service

import (
	"encoding/csv"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	date "github.com/yoshiyoshiharu/item-throw-ways/pkg"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type AreaCollectWeekdayBatchService interface {
  UpdateAll() error
}

type areaCollectWeekdayBatchService struct {
  ar repository.AreaCollectWeekdayRepository
  kr repository.KindRepository
}

func NewAreaCollectWeekdayBatchService(ar repository.AreaCollectWeekdayRepository, kr repository.KindRepository) *areaCollectWeekdayBatchService {
  return &areaCollectWeekdayBatchService{
    ar: ar,
    kr: kr,
  }
}

var (
	AreaCollectWeekdaysApiUrl = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/05syusyubi/syusyubi.csv"
)

type CollectWeekday struct {
	Weekday time.Weekday
	Lap     int
}

func (s *areaCollectWeekdayBatchService) UpdateAll() error {
  allKinds := s.kr.FindAll()

	resp, err := http.Get(
    AreaCollectWeekdaysApiUrl,
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r := csv.NewReader(transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder()))
	rows, err := r.ReadAll()
	if err != nil {
		return err
	}

	var areaCollectWeekdays []entity.AreaCollectWeekday
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

		area := entity.NewArea(area_id, town+street)

		kanenWeekdays := splitWeekday(kanen)
		funenWeekdays := splitNthWeekday(funen)
		shigenWeekday, err := date.JaWeekdayToEn(shigen)
		if err != nil {
			return err
		}

		kindWeekdays := map[string][]CollectWeekday{
			"可燃ごみ": kanenWeekdays,
			"不燃ごみ": funenWeekdays,
			"資源":   {{shigenWeekday, 0}},
		}

		for kindName, weekdays := range kindWeekdays {
			kind := findKind(kindName, allKinds)
			for _, weekday := range weekdays {
				newAreaCollectWeekday := entity.NewAreaCollectWeekday(area, kind, weekday.Weekday, weekday.Lap)
				areaCollectWeekdays = append(areaCollectWeekdays, *newAreaCollectWeekday)
			}
		}
	}

  err = s.ar.DeleteAndInsertAll(areaCollectWeekdays)
  if err != nil {
    return err
  }

  return nil
}

func splitWeekday(str string) []CollectWeekday {
	var collectWeekdays []CollectWeekday
  // 月曜日・金曜日 からweekdayを抜き出す
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
	// 第2・第4火曜日 から lapとweekdayを抜き出す
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

