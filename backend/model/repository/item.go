package repository

import (
	"log"

	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"gorm.io/gorm"
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
  err := Db.Where("name = ?", name).First(&item)

  if err != nil {
		if err.Error != gorm.ErrRecordNotFound {
			log.Fatal(err)
		}
		return false
	}

  return true
}

func (r *itemRepository) GetItems() ([]entity.Item, error) {
  var items []entity.Item
  Db.Find(&items)

  return items, nil
}
