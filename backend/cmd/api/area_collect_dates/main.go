package main

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	area_id, err := strconv.Atoi(request.QueryStringParameters["area_id"])
  if err != nil {
    return events.APIGatewayProxyResponse{}, errors.New("area_id or kind_id is empty")
  }

	var area entity.Area
  err = repository.Db.Where("id = ?", area_id).First(&area).Error

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	repository := repository.NewAreaCollectDatesRepository()
	areaCollectDates := repository.GetAreaCollectDates(area)

	jsonBody, err := json.Marshal(areaCollectDates)

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
