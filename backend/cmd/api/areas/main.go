package main

import (
	"fmt"
	"sort"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

type AreaResponse struct {
  Id string `json:"id"`
  Name string `json:"name"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  repository := repository.NewAreaRepository()
  areas, err := repository.GetAreas()
  sort.Slice(areas, func(i, j int) bool { return areas[i].Id < areas[j].Id })

  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }

	return events.APIGatewayProxyResponse{
    Headers: map[string]string{
      "Access-Control-Allow-Origin": "*",
    },
		Body:       fmt.Sprintf("%v", areas),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
