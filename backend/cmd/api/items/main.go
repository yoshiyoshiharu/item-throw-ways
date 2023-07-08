package main

import (
	"fmt"
	"sort"

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
  items, err := repository.GetItems()
  sort.Slice(items, func(i, j int) bool { return items[i].Id < items[j].Id })

  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }

	return events.APIGatewayProxyResponse{
    Headers: map[string]string{
      "Access-Control-Allow-Origin": "*",
    },
		Body:       fmt.Sprintf("%v", items),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
