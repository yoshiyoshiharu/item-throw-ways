package repository

import (
	"database/sql"
	"log"

	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
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
	sqlStmt := `SELECT name FROM items WHERE name = ?`
	err := Db.QueryRow(sqlStmt, name).Scan(&name)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
		return false
	}

  return true
}

func (r *itemRepository) GetItems() ([]entity.Item, error) {
    rows, err := Db.Query("SELECT id, name FROM items")
  if err != nil {
    return nil, err
  }

  var items []entity.Item
  for rows.Next() {
    var item entity.Item
    err := rows.Scan(&item.Id, &item.Name)
    if err != nil {
      return nil, err
    }
    items = append(items, item)
  }

  return items, nil
}
