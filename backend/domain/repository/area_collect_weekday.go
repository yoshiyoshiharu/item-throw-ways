package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	"gorm.io/gorm"
)

type AreaCollectWeekdayRepository interface {
	FindByAreaId(int) []*entity.AreaCollectWeekday
  DeleteAndInsertAll([]entity.AreaCollectWeekday) error
}

type areaCollectWeekdayRepository struct{
  db *gorm.DB
}

func NewAreaCollectWeekdayRepository(db *gorm.DB) *areaCollectWeekdayRepository {
  return &areaCollectWeekdayRepository{db: db}
}

func (r *areaCollectWeekdayRepository) FindByAreaId(areaId int) []*entity.AreaCollectWeekday {
	var areaCollectWeekdays []*entity.AreaCollectWeekday
	r.db.Joins("Kind").Joins("Area").Where("area_id = ?", areaId).Find(&areaCollectWeekdays)

	return areaCollectWeekdays
}

func (r *areaCollectWeekdayRepository) DeleteAndInsertAll(areaCollectWeekdays []entity.AreaCollectWeekday) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM areas").Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM area_collect_weekdays").Error; err != nil {
			return err
		}
		if err := tx.Omit("Kind").Create(&areaCollectWeekdays).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}
