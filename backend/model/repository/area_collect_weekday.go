package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

type AreaCollectWeekdayRepository interface {
  FindByAreaId(int) []*entity.AreaCollectWeekday
}

type areaCollectWeekdayRepository struct {}

func NewAreaCollectWeekdayRepository() *areaCollectWeekdayRepository {
  return &areaCollectWeekdayRepository{}
}

func (r *areaCollectWeekdayRepository) FindByAreaId(areaId int) []*entity.AreaCollectWeekday {
  var areaCollectWeekdays []*entity.AreaCollectWeekday
  Db.Joins("Kind").Joins("Area").Where("area_id = ?", areaId).Find(&areaCollectWeekdays)

  return areaCollectWeekdays
}
