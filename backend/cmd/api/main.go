package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yoshiyoshiharu/item-throw-ways/di"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/database"
)

func main() {
  router := NewRouter()
  router.Run(":8080")
}

func NewRouter() *gin.Engine {
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

	return router
}
