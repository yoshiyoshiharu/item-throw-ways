package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiyoshiharu/item-throw-ways/model/repository"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

type AreaResponse struct {
  Id string `json:"id"`
  Name string `json:"name"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  repository := repository.NewAreaRepository()
  areas, err := repository.GetAreas()

  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%v", areas),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
