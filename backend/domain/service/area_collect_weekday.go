package service

import (
	"time"

	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	date "github.com/yoshiyoshiharu/item-throw-ways/pkg"
)

type AreaCollectWeekdayService interface {
	ConvertByAreaWithAroundMonths(entity.Area, int, time.Month) []entity.AreaCollectDate
}

type areaCollectWeekdayService struct {
	r repository.AreaCollectWeekdayRepository
}

func NewAreaCollectWeekdayService(repo repository.AreaCollectWeekdayRepository) *areaCollectWeekdayService {
	return &areaCollectWeekdayService{
		r: repo,
	}
}

// error handleing
func (s *areaCollectWeekdayService) ConvertByAreaWithAroundMonths(areaID int, year int, month time.Month) []*entity.AreaCollectDate {
	areaCollectWeekdays := s.r.FindByAreaId(areaID)

	prevYear, prevMonth := date.PrevMonth(year, month)
	nextYear, nextMonth := date.NextMonth(year, month)

	previousMonthAreaCollectDates, err := s.convertFromAreaCollectWeekdays(areaCollectWeekdays, prevYear, prevMonth)
	currentMonthAreaCollectDates, err := s.convertFromAreaCollectWeekdays(areaCollectWeekdays, year, month)
	nextMonthAreaCollectDates, err := s.convertFromAreaCollectWeekdays(areaCollectWeekdays, nextYear, nextMonth)
	if err != nil {
		return nil
	}

	areaCollectDates := []*entity.AreaCollectDate{}
	areaCollectDates = append(areaCollectDates, previousMonthAreaCollectDates...)
	areaCollectDates = append(areaCollectDates, currentMonthAreaCollectDates...)
	areaCollectDates = append(areaCollectDates, nextMonthAreaCollectDates...)

	return areaCollectDates
}

func (s *areaCollectWeekdayService) convertFromAreaCollectWeekdays(areaCollectWeekdays []*entity.AreaCollectWeekday, year int, month time.Month) ([]*entity.AreaCollectDate, error) {
	var areaCollectDates []*entity.AreaCollectDate

	for _, areaCollectWeekday := range areaCollectWeekdays {
		if areaCollectWeekday.Lap == 0 {
			for _, date := range date.AllWeekdayDates(areaCollectWeekday.Weekday, year, month) {
				newAreaCollectDate := entity.NewAreaCollectDate(areaCollectWeekday.Kind, date.Format("2006-01-02"), *&areaCollectWeekday.Area)
				areaCollectDates = append(areaCollectDates, newAreaCollectDate)
			}
		} else {
			date, err := date.NthWeekdayDate(areaCollectWeekday.Lap, areaCollectWeekday.Weekday, year, month)
			if err != nil {
				return nil, err
			}

			newAreaCollectDate := entity.NewAreaCollectDate(areaCollectWeekday.Kind, date.Format("2006-01-02"), *&areaCollectWeekday.Area)
			areaCollectDates = append(areaCollectDates, newAreaCollectDate)
		}
	}

	return areaCollectDates, nil
}
