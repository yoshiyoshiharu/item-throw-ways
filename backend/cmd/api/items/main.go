package main

import (
	"encoding/json"
	"sort"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  repository := repository.NewItemRepository()
  items, err := repository.GetItems()
  sort.Slice(items, func(i, j int) bool { return items[i].ID < items[j].ID })

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
