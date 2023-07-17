package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll(int) []*entity.Item
}

type itemRepository struct{}

func NewItemRepository() *itemRepository {
	return &itemRepository{}
}

func (r *itemRepository) FindAll() []*entity.Item {
	var items []*entity.Item
	Db.Preload("Kinds").Find(&items)

	return items
}

func (r *itemRepository) DeleteAndInsertAll(items []*entity.Item) error {
	err := Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM items").Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM item_kinds").Error; err != nil {
			return err
		}
		if err := tx.Create(&items).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}
