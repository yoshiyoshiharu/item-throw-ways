package repository

import (
	"errors"

	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

type KindRepository interface {
  GetKinds() []entity.Kind
  GetKindIdByName([]entity.Kind, string) (int, error)
}

type kindRepository struct {}

func NewKindRepository() KindRepository {
	return &kindRepository{}
}

func (r *kindRepository) GetKinds() []entity.Kind {
  kinds := []entity.Kind{}

  Db.Find(&kinds)
  return kinds
}

func (r *kindRepository) GetKindIdByName(kinds []entity.Kind, name string) (int, error) {
  for _, kind := range kinds {
    if kind.Name == name {
      return kind.Id, nil
    }
  }

  return 0, errors.New("Not found")
}

