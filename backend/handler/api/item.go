package handler

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/api"
)

type ItemHandler interface {
	FindAll(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type itemHandler struct {
	s service.ItemService
}

func NewItemHandler(service service.ItemService) *itemHandler {
	return &itemHandler{
		s: service,
	}
}

func (h *itemHandler) FindAll(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	items := h.s.FindAll()

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
