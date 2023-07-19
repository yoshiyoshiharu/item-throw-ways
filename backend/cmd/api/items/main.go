package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	"github.com/yoshiyoshiharu/item-throw-ways/domain/service"
	"github.com/yoshiyoshiharu/item-throw-ways/handler"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/database"
)

func main() {
  db, err := database.Connect()
  if err != nil {
    panic(err)
  }

  r := repository.NewItemRepository(db)
  s := service.NewItemService(r)
  h := handler.NewItemHandler(s)

	lambda.Start(h.FindAll)
}
