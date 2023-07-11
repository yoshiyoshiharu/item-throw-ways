package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

type ItemRepository interface {
  GetItems() ([]entity.Item, error)
  ItemExists(string) bool
}

type itemRepository struct {}

func NewItemRepository() ItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) ItemExists(name string) bool {
  var item entity.Item
  result := Db.Where("name = ?", name).Limit(1).Find(&item)

  if result.RowsAffected == 0 {
    return false
  }

  return true
}

func (r *itemRepository) GetItems() ([]entity.Item, error) {
  var items []entity.Item
  Db.Preload("Kinds").Find(&items)

  return items, nil
}
