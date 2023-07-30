package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	di "github.com/yoshiyoshiharu/item-throw-ways/di/batch"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

  h := di.InitAreaCollectWeekday(db)

  lambda.Start(h.UpdateAll)
}
