package repository

import (
	"errors"
	"log"

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

  rows, err := Db.Query("SELECT id, name FROM kinds;")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    var kind entity.Kind
    err := rows.Scan(&kind.Id, &kind.Name)
    if err != nil {
      log.Fatal(err)
    }
    kinds = append(kinds, kind)
  }

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

