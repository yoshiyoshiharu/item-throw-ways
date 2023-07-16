package repository

import (
	"time"

	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	date "github.com/yoshiyoshiharu/item-throw-ways/pkg"
)

type AreaCollectDatesRepository interface {
  GetAreaCollectDates(area entity.Area) []entity.AreaCollectDate
}

type areaCollectDatesRepository struct {}

func NewAreaCollectDatesRepository() *areaCollectDatesRepository {
  return &areaCollectDatesRepository{}
}

func (r *areaCollectDatesRepository) GetAreaCollectDates(area entity.Area, year int, month time.Month) []entity.AreaCollectDate {
  var areaCollectWeekdays []entity.AreaCollectWeekday
  var areaCollectDates []entity.AreaCollectDate

  Db.Joins("Kind").Where("area_id = ?", area.ID).Find(&areaCollectWeekdays)

  for _, areaCollectWeekday := range areaCollectWeekdays {
    if areaCollectWeekday.Lap == 0 {
      for _, date := range date.AllWeekdayDates(areaCollectWeekday.Weekday, year, month) {
        insertDate := entity.AreaCollectDate{Kind: areaCollectWeekday.Kind.Name, Date: date.Format("2006-01-02")}
        areaCollectDates = append(areaCollectDates, insertDate)
      }
    } else {
      date, err := date.NthWeekdayDate(areaCollectWeekday.Lap, areaCollectWeekday.Weekday, year, month)
      if err != nil {
        panic(err)
      }

      insertDate := entity.AreaCollectDate{Kind: areaCollectWeekday.Kind.Name, Date: date.Format("2006-01-02")}
      areaCollectDates = append(areaCollectDates, insertDate)
    }
  }

  return areaCollectDates
}
