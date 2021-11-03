package service

import (
	"encoding/json"

	"github.com/bburaksseyhan/orderapi/src/cmd/utils"
	"github.com/bburaksseyhan/orderapi/src/pkg/model"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

const (
	CREATEQUEUENAME    = "NEW_ORDER_QUEUE"
	CANCELLEDQUEUENAME = "CANCELLED_ORDER_QUEUE"
)

func CreateOrder(setting *utils.QueueSettings, order model.Order) error {
	return queueService(CREATEQUEUENAME, setting, order)
}

func CancelledOrder(setting *utils.QueueSettings, order model.Order) error {
	return queueService(CANCELLEDQUEUENAME, setting, order)
}

func queueService(queueName string, setting *utils.QueueSettings, order model.Order) error {
	conn, err := amqp.Dial(setting.Url)

	if err != nil {
		log.Error("Failed Initializing Broker Connection")
		panic(err.Error())
	}

	channel, err := conn.Channel()

	if err != nil {
		log.Error(err)
	}

	defer channel.Close()

	q, _ := channel.QueueDeclare(queueName, false, false, false, false, nil)

	log.Info(q)

	//convert order to byte array
	marshal, _ := json.Marshal(order)

	if err = channel.Publish("", queueName, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        marshal,
		},
	); err != nil {
		log.Error(err)
	}

	return err
}
