package date

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNthWeekdayDate(t *testing.T) {
  tests := []struct {
    lap int
    weekday string
    year int
    month time.Month
    want time.Time
  }{
    {1, "月曜日", 2023, 7, time.Date(2023, 7, 3, 0, 0, 0, 0, loc)},
    {3, "木曜日", 2023, 7, time.Date(2023, 7, 20, 0, 0, 0, 0, loc)},
    {4, "土曜日", 2023, 6, time.Date(2023, 6, 24, 0, 0, 0, 0, loc)},
    {2, "土曜日", 2023, 8, time.Date(2023, 8, 12, 0, 0, 0, 0, loc)},
    {1, "土曜日", 2023, 7, time.Date(2023, 7, 1, 0, 0, 0, 0, loc)},
    {5, "月曜日", 2023, 7, time.Date(2023, 7, 31, 0, 0, 0, 0, loc)},
  }

  for _, test := range tests {
    t.Run("[正常系]" + strconv.Itoa(test.year) + "年" + strconv.Itoa(int(test.month)) + "月の第" + strconv.Itoa(test.lap) + test.weekday + "の日付を返す", func(t *testing.T) {
      weekday, _ := JaWeekdayToEn(test.weekday)
      result, err := NthWeekdayDate(test.lap, weekday, test.year, test.month)

      assert.NoError(t, err)
      assert.Equal(t, test.want, result)
    })
  }

  t.Run("[異常系]存在しない日付になるとエラーを返す", func(t *testing.T) {
    _, err := NthWeekdayDate(6, 6, 2023, 7)
    assert.Error(t, err)
  })
}
