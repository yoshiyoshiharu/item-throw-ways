package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/pkg/database"
)

type AreaRepository interface {
  FindAll() []*entity.Area
  FindById(areaId int) (*entity.Area, error)
}

type areaRepository struct {}

func NewAreaRepository() *areaRepository {
  return &areaRepository{}
}

func (r *areaRepository) FindAll() []entity.Area {
  var areas []entity.Area
  database.Db.Find(&areas)

  return areas
}

func (r *areaRepository) FindById(id int) (entity.Area, error) {
  var area entity.Area
  err := database.Db.Where("id = ?", id).First(&area).Error
  if err != nil {
    return entity.Area{}, err
  }

  return area, nil
}
