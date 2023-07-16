package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/pkg/database"
)

type AreaRepository interface {
  GetAreas() ([]entity.Area, error)
}

type areaRepository struct {}

func NewAreaRepository() AreaRepository {
  return &areaRepository{}
}

func (r *areaRepository) GetAreas() ([]entity.Area, error) {
  var areas []entity.Area
  database.Db.Find(&areas)

  return areas, nil
}
