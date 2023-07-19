package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/yoshiyoshiharu/item-throw-ways/domain/service"
)

type AreaCollectDateHandler interface {
  FindByAreaId(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type areaCollectDateHandler struct {
  s service.AreaCollectWeekdayService
}

func NewAreaCollectDateHandler(service service.AreaCollectWeekdayService) *areaCollectDateHandler {
  return &areaCollectDateHandler {
    s: service,
  }
}

func (h *areaCollectDateHandler) FindAll (request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	area_id, year, month, err := parseParams(request)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errors.New("request parameter is invalid")
	}

	areas := h.s.ConvertByAreaWithAroundMonths(area_id, year, month)

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

func parseParams(request events.APIGatewayProxyRequest) (int, int, time.Month, error) {
	area_id, err := strconv.Atoi(request.QueryStringParameters["area_id"])
	year, err := strconv.Atoi(request.QueryStringParameters["year"])
	monthInt, err := strconv.Atoi(request.QueryStringParameters["month"])
	month, err := intToMonth(monthInt)

	if err != nil {
		return 0, 0, 0, err
	}
	return area_id, year, month, nil
}

func intToMonth(monthInt int) (time.Month, error) {
	if monthInt < 1 || monthInt > 12 {
		return 0, fmt.Errorf("Invalid month: %d", monthInt)
	}
	return time.Month(monthInt), nil
}
