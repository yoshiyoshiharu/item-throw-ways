package date

import (
	"errors"
	"fmt"
	"time"
)

var loc = time.FixedZone("Asia/Tokyo", 9*60*60)

var JaWeekdays = map[string]string{
	"Sunday":    "日曜日",
	"Monday":    "月曜日",
	"Tuesday":   "火曜日",
	"Wednesday": "水曜日",
	"Thursday":  "木曜日",
	"Friday":    "金曜日",
	"Saturday":  "土曜日",
}

func NthWeekdayDate(lap int, wd time.Weekday, year int, month time.Month) (time.Time, error) {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)

	firstWeekday := firstDay.AddDate(0, 0, int(wd-firstDay.Weekday()))

	if firstWeekday.Month() != firstDay.Month() {
		firstWeekday = firstWeekday.AddDate(0, 0, 7)
	}

	nthWeekday := firstWeekday.AddDate(0, 0, 7*(lap-1))

	if nthWeekday.Month() != firstDay.Month() {
		err := fmt.Errorf("%d年%d月の第%d %sは存在しません。", year, month, lap, wd)
		return time.Time{}, err
	}

	return nthWeekday, nil
}

func AllWeekdayDates(wd time.Weekday, year int, month time.Month) []time.Time {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)

	firstWeekday := firstDay.AddDate(0, 0, int(wd-firstDay.Weekday()))

	if firstWeekday.Month() != firstDay.Month() {
		firstWeekday = firstWeekday.AddDate(0, 0, 7)
	}

	var dates []time.Time
	for i := 0; i < 5; i++ {
		if firstWeekday.AddDate(0, 0, 7*i).Month() == firstDay.Month() {
			dates = append(dates, firstWeekday.AddDate(0, 0, 7*i))
		}
	}

	return dates
}

func JaWeekdayToEn(ja string) (time.Weekday, error) {
	switch ja {
	case JaWeekdays["Sunday"]:
		return time.Sunday, nil
	case JaWeekdays["Monday"]:
		return time.Monday, nil
	case JaWeekdays["Tuesday"]:
		return time.Tuesday, nil
	case JaWeekdays["Wednesday"]:
		return time.Wednesday, nil
	case JaWeekdays["Thursday"]:
		return time.Thursday, nil
	case JaWeekdays["Friday"]:
		return time.Friday, nil
	case JaWeekdays["Saturday"]:
		return time.Saturday, nil
	default:
		return time.Sunday, errors.New("曜日の変換に失敗しました。")
	}
}

func PrevMonth(year int, month time.Month) (int, time.Month) {
	if month == 1 {
		return year - 1, 12
	}
	return year, month - 1
}

func NextMonth(year int, month time.Month) (int, time.Month) {
	if month == 12 {
		return year + 1, 1
	}
	return year, month + 1
}
