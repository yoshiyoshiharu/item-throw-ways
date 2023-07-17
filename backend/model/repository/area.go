package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

type AreaRepository interface {
  FindAll() []*entity.Area
  FindById(int) (*entity.Area, error)
}

type areaRepository struct {}

func NewAreaRepository() *areaRepository {
  return &areaRepository{}
}

func (r *areaRepository) FindAll() []*entity.Area {
  var areas []*entity.Area
  Db.Find(&areas)

  return areas
}

func (r *areaRepository) FindById(id int) (*entity.Area, error) {
  var area *entity.Area
  err := Db.Where("id = ?", id).First(&area).Error
  if err != nil {
    return nil, err
  }

  return area, nil
}
