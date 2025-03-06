package infraestructure

import (
	"apihex01/core"
	"apihex01/students/domain/entities"
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	conn *core.Conn_Rabbit
}

func NewRabbitMq() *Rabbit {
	conn := core.GetConnRabbit()

	if conn.Err != "" {
		log.Fatalf("Error al tratar de hacer una conexión hacia rabbit: %v", conn.Err)
	}

	return &Rabbit{conn: conn}

}

func (r *Rabbit) SendMessageToBroker(student *entities.Student) {

	defer r.conn.Broker.Close()
	defer r.conn.Channel.Close()

	err := r.conn.Channel.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	r.conn.FailOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(student)
	r.conn.FailOnError(err, "Failed to marshal JSON")

	err = r.conn.Channel.PublishWithContext(ctx,
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	r.conn.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)

}
