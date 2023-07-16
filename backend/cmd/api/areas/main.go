package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/pkg/database"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  var areas []entity.Area
  err := database.Db.Find(&areas).Order("id").Error
  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }

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
