package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

type ItemRepository interface {
  FindAll(int) []*entity.Item
}

type itemRepository struct {}

func NewItemRepository() *itemRepository {
  return &itemRepository{}
}

func (r *itemRepository) FindAll() []*entity.Item {
  var items []*entity.Item
  Db.Preload("Kinds").Find(&items)

  return items
}
