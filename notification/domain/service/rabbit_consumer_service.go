package service

import (
	"heptaber/notification/domain/model"
	"heptaber/notification/infrastructure/rabbitmq"
	"log"
	"strings"
)

func ConsumeVerificationEmailNotifications() {
	rmq, err := rabbitmq.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err.Error())
		return
	}
	defer rmq.CloseConnection()

	rmq.SetUpForNotification()

	msgs, err := rmq.ConsumeMessages(rabbitmq.NotificationQueueName)
	if err != nil {
		log.Fatalf("Failed to consume messages: %s", err.Error())
		return
	}

	for msg := range msgs {
		routingKey := msg.RoutingKey
		msgAr := strings.Split(routingKey, ".")
		var msgEvent string = msgAr[0]
		var msgType string = msgAr[1]

		switch msgType {
		case string(model.EMAIL):
			SendEmail(msgEvent, []byte(msg.Body))
		case string(model.NOTIFICATION):
			// notificationService.SendNotification(msgEvent, blah-blah)
		default:
			log.Fatalf("Unknown message type: %s", msgType)
		}
	}
}
