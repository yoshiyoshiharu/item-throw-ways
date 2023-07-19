package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/database"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  db, err := database.Connect()
  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }

	areaRepository := repository.NewAreaRepository(db)
	areas := areaRepository.FindAll()

	jsonBody, err := json.Marshal(areas)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
		Body:       string(jsonBody),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
