package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/model/database"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  db, err := database.Connect()
  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }

	itemRepository := repository.NewItemRepository(db)
	items := itemRepository.FindAll()

	jsonBody, err := json.Marshal(items)

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
