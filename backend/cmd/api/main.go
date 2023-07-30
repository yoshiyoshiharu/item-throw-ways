package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	di "github.com/yoshiyoshiharu/item-throw-ways/di/api"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/database"
)

var ginLambda *ginadapter.GinLambda
func init() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowCredentials: false,
	}))

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

  areaHandler := di.InitArea(db)
  itemHandler := di.InitItem(db)
  areaCollectDateHandler := di.InitAreaCollectDate(db)

	router.GET("/areas", areaHandler.FindAll)
  router.GET("/items", itemHandler.FindAll)
  router.GET("/area_collect_dates", areaCollectDateHandler.FindAll)

  ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
  lambda.Start(Handler)
}
