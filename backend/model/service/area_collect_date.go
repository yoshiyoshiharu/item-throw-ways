package service

import (
	"time"

	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
	date "github.com/yoshiyoshiharu/item-throw-ways/pkg"
)

type AreaCollectDateService interface {
	GetByAreaWithAroundMonths(entity.Area, int, time.Month) []entity.AreaCollectDate
}

type areaCollectDateService struct {
	r repository.AreaCollectWeekdayRepository
}

func NewAreaCollectDateService(repo repository.AreaCollectWeekdayRepository) *areaCollectDateService {
	return &areaCollectDateService{
		r: repo,
	}
}

func (s *areaCollectDateService) GetByAreaWithAroundMonths(area *entity.Area, year int, month time.Month) []*entity.AreaCollectDate {
	areaCollectWeekdays := s.r.FindByAreaId(area.ID)

	prevYear, prevMonth := date.PrevMonth(year, month)
	nextYear, nextMonth := date.NextMonth(year, month)

	previousMonthAreaCollectDates, err := s.getByAreaCollectWeekdays(areaCollectWeekdays, prevYear, prevMonth)
	currentMonthAreaCollectDates, err := s.getByAreaCollectWeekdays(areaCollectWeekdays, year, month)
	nextMonthAreaCollectDates, err := s.getByAreaCollectWeekdays(areaCollectWeekdays, nextYear, nextMonth)
	if err != nil {
		return nil
	}

	areaCollectDates := []*entity.AreaCollectDate{}
	areaCollectDates = append(areaCollectDates, previousMonthAreaCollectDates...)
	areaCollectDates = append(areaCollectDates, currentMonthAreaCollectDates...)
	areaCollectDates = append(areaCollectDates, nextMonthAreaCollectDates...)

	return areaCollectDates
}

func (s *areaCollectDateService) getByAreaCollectWeekdays(areaCollectWeekdays []*entity.AreaCollectWeekday, year int, month time.Month) ([]*entity.AreaCollectDate, error) {
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
