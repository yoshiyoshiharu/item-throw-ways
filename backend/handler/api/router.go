package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

  areaHandler := NewAreaHandler()
	router.GET("/areas", areaHandler.s.FindAll())

	return router
}
