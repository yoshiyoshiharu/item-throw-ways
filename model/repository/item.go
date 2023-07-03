package repository

import (
	"database/sql"
	"log"
)

type ItemRepository interface {
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
