package rabbitmq

import (
	"heptaber/auth/app/initializers"
	"log"
	"os"

	"github.com/streadway/amqp"
)

const (
	NotificationExchangeName    string = "notification"
	VerificaitonEmailRoutingKey string = "verification.email"
	NotificationQueueName       string = "notification_queue"
)

type RabbitMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewRabbitMQ() (*RabbitMQ, error) {
	initializers.LoadEnvVariables()
	rmqUser := os.Getenv("RABBITMQ_USER")
	rmqPassword := os.Getenv("RABBITMQ_PASSWORD")
	rmqHost := os.Getenv("RABBITMQ_HOST")
	rmqPort := os.Getenv("RABBITMQ_PORT")
	conn, err := amqp.Dial(
		"amqp://" +
			rmqUser + ":" + rmqPassword + "@" +
			rmqHost + ":" + rmqPort + "/",
	)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	rabbitMQ := &RabbitMQ{
		connection: conn,
		channel:    ch,
	}

	return rabbitMQ, nil
}

func (r *RabbitMQ) PublishMessage(exchange, routingKey string, body []byte) error {
	err := r.channel.Publish(
		exchange,   // exchange name
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) DeclareQueue(queueName string) error {
	queue, err := r.channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // autodelete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	r.queue = queue

	return nil
}

func (r *RabbitMQ) ExchangeDeclare(exchangeName string) (err error) {
	err = r.channel.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	return
}

func (r *RabbitMQ) QueueBind(queueName, routingKey, exchangeName string) (err error) {
	err = r.channel.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	return
}

func (r *RabbitMQ) ConsumeMessages(queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := r.channel.Consume(
		queueName, // queue name очереди
		"",        // consumer name
		true,      // autodelete messages
		false,     // don't use unique identifiers
		false,     // don't use preview
		false,     // block consume
		nil,       // extra args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (r *RabbitMQ) CloseConnection() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.connection != nil {
		r.connection.Close()
	}
}

func (r *RabbitMQ) SetUpForNotification() {
	err := r.ExchangeDeclare(NotificationExchangeName)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %s", err.Error())
		return
	}

	err = r.DeclareQueue(NotificationQueueName)
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err.Error())
	}

	err = r.QueueBind(NotificationQueueName, VerificaitonEmailRoutingKey, NotificationExchangeName)
	if err != nil {
		log.Fatalf("Failed to bind queue: %s", err.Error())
	}
}
