package repository

import (
	"fmt"

	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"gorm.io/gorm"
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

func (r *areaCollectWeekdayRepository) DeleteAndInsertAll(areaCollectWeekdays []*entity.AreaCollectWeekday) error {
  fmt.Println(areaCollectWeekdays)
  err := Db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Exec("DELETE FROM areas").Error; err != nil {
      return err
    }
    if err := tx.Exec("DELETE FROM area_collect_weekdays").Error; err != nil {
      return err
    }
    if err := tx.Create(&areaCollectWeekdays).Error; err != nil {
      return err
    }
    return nil
  })

  return err
}
