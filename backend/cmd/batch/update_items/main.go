package main

import (
	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/batch"
	handler "github.com/yoshiyoshiharu/item-throw-ways/handler/batch"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	ir := repository.NewItemRepository(db)
	kr := repository.NewKindRepository(db)
	s := service.NewItemBatchService(ir, kr)
	h := handler.NewItemBatchHandler(s)

  h.UpdateAll()
}
