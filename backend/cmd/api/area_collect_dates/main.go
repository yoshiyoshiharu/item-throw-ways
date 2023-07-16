package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	area_id, err := strconv.Atoi(request.QueryStringParameters["area_id"])
	year, err := strconv.Atoi(request.QueryStringParameters["year"])
	monthInt, err := strconv.Atoi(request.QueryStringParameters["month"])
  month, err := intToMonth(monthInt)
  if err != nil {
    return events.APIGatewayProxyResponse{}, errors.New("request parameter is invalid")
  }

	var area entity.Area
  err = repository.Db.Where("id = ?", area_id).First(&area).Error

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	repository := repository.NewAreaCollectDatesRepository()
	areaCollectDates := repository.GetAreaCollectDatesWithAroundMonthes(area, year, month)

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

func intToMonth(monthInt int) (time.Month, error) {
	if monthInt < 1 || monthInt > 12 {
		return 0, fmt.Errorf("Invalid month: %d", monthInt)
	}
	return time.Month(monthInt), nil
}

