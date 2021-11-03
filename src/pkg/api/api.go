package api

import (
	"fmt"

	"github.com/bburaksseyhan/orderapi/src/cmd/utils"
	handlers "github.com/bburaksseyhan/orderapi/src/pkg/handler"
	"github.com/gin-gonic/gin"
)

func Initialize(config utils.Configuration) {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	fmt.Printf("%+v\n", config)

	orderHandler := handlers.NewOrderHandler(&config.Queue)

	router.POST("/", orderHandler.CreateOrder)
	router.POST("/cancelled", orderHandler.CancelledOrder)

	router.GET("/ping", orderHandler.HealthCheck)

	router.Run(":" + config.Server.Port + "")
}
