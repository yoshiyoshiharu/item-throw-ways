package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

type AreaCollectWeekdaysRepository interface {
	GetAreaCollectWeekdays(entity.Area, entity.Kind)
}

type areaCollectWeekdaysRepository struct{}

func NewAreaCollectWeekdaysRepository() *areaCollectWeekdaysRepository {
	return &areaCollectWeekdaysRepository{}
}

func (repository *areaCollectWeekdaysRepository) GetAreaCollectWeekdays(area entity.Area, kind entity.Kind) []entity.AreaCollectWeekday {
	var areaCollectWeekdays []entity.AreaCollectWeekday
	Db.Where("area_id = ? AND kind_id = ?", area.Id, kind.Id).Find(&areaCollectWeekdays)

	return areaCollectWeekdays
}
