package service

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

type KindService interface {
  FindAll() []*entity.Kind
}

type kindService struct {
  r repository.KindRepository
}

func NewKindService(repo repository.KindRepository) *kindService {
  return &kindService{
    r: repo,
  }
}

func (s *kindService) FindAll() []*entity.Kind {
  return s.r.FindAll()
}

