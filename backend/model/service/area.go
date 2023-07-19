package service

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

type AreaService interface {
  FindAll() []*entity.Area
  FindById(int) (*entity.Area, error)
}

type areaService struct {
  r repository.AreaRepository
}

func NewAreaService(repo repository.AreaRepository) *areaService {
  return &areaService{
    r: repo,
  }
}

func (s *areaService) FindAll() []*entity.Area {
  return s.r.FindAll()
}

func (s *areaService) FindById(id int) (*entity.Area, error) {
  area, err := s.r.FindById(id)

  if err != nil {
    return nil, err
  }

  return area, nil
}
