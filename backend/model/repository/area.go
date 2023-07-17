package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"gorm.io/gorm"
)

type AreaRepository interface {
	FindAll() []*entity.Area
	FindById(int) (*entity.Area, error)
}

type areaRepository struct{
  db *gorm.DB
}

func NewAreaRepository(db *gorm.DB) *areaRepository {
	return &areaRepository{db: db}
}

func (r *areaRepository) FindAll() []*entity.Area {
	var areas []*entity.Area
	r.db.Find(&areas)

	return areas
}

func (r *areaRepository) FindById(id int) (*entity.Area, error) {
	var area *entity.Area
	err := r.db.Where("id = ?", id).First(&area).Error
	if err != nil {
		return nil, err
	}

	return area, nil
}
