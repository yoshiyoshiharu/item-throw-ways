package handler

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/yoshiyoshiharu/item-throw-ways/domain/service"
)

type AreaHandler interface {
  FindAll() (events.APIGatewayProxyResponse, error)
}

type areaHandler struct {
  s service.AreaService
}

func NewAreaHandler(service service.AreaService) *areaHandler {
  return &areaHandler{
    s: service,
  }
}

func (h *areaHandler)FindAll (request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	areas := h.s.FindAll()

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

