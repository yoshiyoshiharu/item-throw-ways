package repository

import "github.com/yoshiyoshiharu/item-throw-ways/model/entity"

type AreaRepository interface {
  GetAreas() ([]entity.Area, error)
}

type areaRepository struct {}

func NewAreaRepository() AreaRepository {
  return &areaRepository{}
}

func (r *areaRepository) GetAreas() ([]entity.Area, error) {
  rows, err := Db.Query("SELECT id, name FROM areas")
  if err != nil {
    return nil, err
  }

  var areas []entity.Area
  for rows.Next() {
    var area entity.Area
    err := rows.Scan(&area.Id, &area.Name)
    if err != nil {
      return nil, err
    }
    areas = append(areas, area)
  }

  return areas, nil
}
