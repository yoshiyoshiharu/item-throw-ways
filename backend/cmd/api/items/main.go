package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

type ItemResponse struct {
  Id          string `json:"id"`
  Name        string `json:"name"`
  Kind        string `json:"kind"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  repository := repository.NewItemRepository()
  areas, err := repository.GetItems()

  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%v", areas),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
