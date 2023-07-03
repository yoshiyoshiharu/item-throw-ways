package date

import (
	"fmt"
	"time"
)

var loc = time.FixedZone("Asia/Tokyo", 9*60*60)

func NthWeekdayDate(n int, wd time.Weekday, year int, month int) (time.Time, error) {
  firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)

  firstWeekday := firstDay.AddDate(0, 0, int(wd - firstDay.Weekday()))

  if firstWeekday.Month() != firstDay.Month() {
    firstWeekday = firstWeekday.AddDate(0, 0, 7)
  }

  nthWeekday := firstWeekday.AddDate(0, 0, 7 * (n - 1))

  if nthWeekday.Month() != firstDay.Month() {
    err := fmt.Errorf("%d年%d月の第%d %sは存在しません。", year, month, n, wd)
    return time.Time{}, err
  }

  return nthWeekday, nil
}
