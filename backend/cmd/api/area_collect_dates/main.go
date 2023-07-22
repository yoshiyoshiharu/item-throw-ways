package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	"github.com/yoshiyoshiharu/item-throw-ways/domain/service"
	handler "github.com/yoshiyoshiharu/item-throw-ways/handler/api"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/database"
)

func main() {
  db, err := database.Connect()
  if err != nil {
    panic(err)
  }

  r := repository.NewAreaCollectWeekdayRepository(db)
  s := service.NewAreaCollectWeekdayService(r)
  h := handler.NewAreaCollectDateHandler(s)

	lambda.Start(h.FindAll)
}

