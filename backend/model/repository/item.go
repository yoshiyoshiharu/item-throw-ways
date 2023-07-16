package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/pkg/database"
)

type ItemRepository interface {
  FindAll(int) []*entity.Item
}

type itemRepository struct {}

func NewItemsRepository() *itemRepository {
  return &itemRepository{}
}

func (r *itemRepository) FindAll() []*entity.Item {
  var items []*entity.Item
  database.Db.Preload("Kinds").Find(&items)

  return items
}
