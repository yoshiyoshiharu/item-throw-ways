package main

import (
	"encoding/json"
	"sort"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  repository := repository.NewAreaRepository()
  areas, err := repository.GetAreas()
  sort.Slice(areas, func(i, j int) bool { return areas[i].Id < areas[j].Id })

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
