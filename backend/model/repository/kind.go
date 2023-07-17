package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"gorm.io/gorm"
)

type KindRepository interface {
	FindAll(int) []*entity.Kind
}

type kindRepository struct{
  db *gorm.DB
}

func NewKindRepository(db *gorm.DB) *kindRepository {
  return &kindRepository{db: db}
}

func (r *kindRepository) FindAll() []*entity.Kind {
	var kinds []*entity.Kind
	r.db.Preload("Kinds").Find(&kinds)

	return kinds
}
