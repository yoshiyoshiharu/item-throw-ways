package service

import (
	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
)

type ItemService interface {
  FindAll() []*entity.Item
  DeleteAndInsertAll([]*entity.Item) error
}

type itemService struct {
  r repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) *itemService {
  return &itemService{
    r: repo,
  }
}

func (s *itemService) FindAll() []*entity.Item {
  return s.r.FindAll()
}


func (s *itemService) DeleteAndInsertAll(items []*entity.Item) error {
  return s.r.DeleteAndInsertAll(items)
}
