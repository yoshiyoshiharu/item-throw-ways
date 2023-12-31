package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll() []*entity.Item
	DeleteAndInsertAll([]entity.Item) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) FindAll() []*entity.Item {
	var items []*entity.Item
	r.db.Preload("Kinds").Find(&items)

	return items
}

func (r *itemRepository) DeleteAndInsertAll(items []entity.Item) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
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
