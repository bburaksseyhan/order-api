package handler

import (
	"net/http"
	"os"

	"github.com/bburaksseyhan/orderapi/src/cmd/utils"
	"github.com/bburaksseyhan/orderapi/src/pkg/model"
	createOrderService "github.com/bburaksseyhan/orderapi/src/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type OrderHandler interface {
	CreateOrder(*gin.Context)
	CancelledOrder(*gin.Context)

	HealthCheck(*gin.Context)
}

type orderHandler struct {
	settings utils.QueueSettings
}

func NewOrderHandler(settings *utils.QueueSettings) OrderHandler {
	return &orderHandler{
		settings: *settings,
	}
}

func (o *orderHandler) CreateOrder(c *gin.Context) {
	var order model.Order

	c.BindJSON(&order)

	err := createOrderService.CreateOrder(&o.settings, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (o *orderHandler) CancelledOrder(c *gin.Context) {
	var order model.Order

	c.BindJSON(&order)

	err := createOrderService.CancelledOrder(&o.settings, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (o *orderHandler) HealthCheck(c *gin.Context) {

	log.Info("ping == pong")

	hostname, err := os.Hostname()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info("hostname", hostname)

	c.JSON(http.StatusOK, gin.H{"ping": hostname})
}
