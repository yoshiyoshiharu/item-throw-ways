package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/model/entity"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

type AreaCollectingDatesResponse struct {
	AreaCollectWeekdays []entity.AreaCollectWeekday `json:"area_collect_weekdays"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	area_id := request.QueryStringParameters["area_id"]
	kind_id := request.QueryStringParameters["kind_id"]

	var kind entity.Kind
	var area entity.Area
	err := repository.Db.Where("id = ?", kind_id).First(&kind).Error
	err = repository.Db.Where("id = ?", area_id).First(&area).Error

	if err != nil {
		log.Fatal(err)
		return events.APIGatewayProxyResponse{}, err
	}

	repository := repository.NewAreaCollectWeekdaysRepository()
	areaCollectWeekdays := repository.GetAreaCollectWeekdays(area, kind)

	jsonBody, err := json.Marshal(areaCollectWeekdays)

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
