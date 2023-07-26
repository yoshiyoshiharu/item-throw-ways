package main

import (
	di "github.com/yoshiyoshiharu/item-throw-ways/di/batch"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

  h := di.InitItemBatch(db)

  h.UpdateAll()
}
