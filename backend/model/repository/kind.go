package repository

import (
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

type KindRepository interface {
	FindAll(int) []*entity.Kind
}

type kindRepository struct{}

func NewKindRepository() *kindRepository {
	return &kindRepository{}
}

func (r *kindRepository) FindAll() []*entity.Kind {
	var kinds []*entity.Kind
	Db.Preload("Kinds").Find(&kinds)

	return kinds
}
