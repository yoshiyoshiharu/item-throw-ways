package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/pkg/database"
)

type AreaCollectWeekdayRepository interface {
  FindByAreaId(areaId int) []*entity.AreaCollectWeekday
}

type areaCollectWeekdayRepository struct {}

func NewAreaCollectWeekdaysRepository() *areaCollectWeekdayRepository {
  return &areaCollectWeekdayRepository{}
}

func (r *areaCollectWeekdayRepository) FindByAreaId(areaId int) []entity.AreaCollectWeekday {
  var areaCollectWeekdays []entity.AreaCollectWeekday
  database.Db.Joins("Kind").Joins("Area").Where("area_id = ?", areaId).Find(&areaCollectWeekdays)

  return areaCollectWeekdays
}
