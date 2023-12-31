package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	"gorm.io/gorm"
)

type KindRepository interface {
	FindAll() []*entity.Kind
}

type kindRepository struct {
	db *gorm.DB
}

func NewKindRepository(db *gorm.DB) KindRepository {
	return &kindRepository{db: db}
}

func (r *kindRepository) FindAll() []*entity.Kind {
	var kinds []*entity.Kind
	r.db.Find(&kinds)

	return kinds
}
