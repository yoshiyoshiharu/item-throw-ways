package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

type ItemRepository interface {
  GetItems() ([]entity.Item, error)
}

type itemRepository struct {}

func NewItemRepository() ItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) GetItems() ([]entity.Item, error) {
  var items []entity.Item
  Db.Preload("Kinds").Find(&items)

  return items, nil
}
