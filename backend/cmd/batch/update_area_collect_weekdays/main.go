package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/batch"
	handler "github.com/yoshiyoshiharu/item-throw-ways/handler/batch"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/database"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
)

var (
	areas    []entity.Area
	allKinds []*entity.Kind
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	ar := repository.NewAreaCollectWeekdayRepository(db)
	kr := repository.NewKindRepository(db)
	s := service.NewAreaCollectWeekdayBatchService(ar, kr)
	h := handler.NewAreaCollectWeekdayBatchHandler(s)

	lambda.Start(h.UpdateAll)
}
