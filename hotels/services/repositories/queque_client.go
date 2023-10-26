package repositories

import (
	"context"
	"fmt"
	e "hotels/utils/errors"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Si ocurre un error muestra un mensaje y se termina el programa
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// Esta estructura se utiliza para interactuar con el servidor RabbitMQ.
type QueueClient struct {
	Connection *amqp.Connection
}

// crea una conexión al servidor RabbitMQ
func NewQuequeClient(user string, pass string, host string, port int) *QueueClient {
	Connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", user, pass, host, port))
	failOnError(err, "Failed to connect to RabbitMQ")
	return &QueueClient{
		Connection: Connection,
	}
}

// Detalles del hotel
type HotelMessage struct {
	HotelId     string  `json:"hotel_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Amenities   string  `json:"amenities"`
	Stars       string  `json:"stars"`
	Rooms       int     `json:"rooms"`
	Price       float32 `json:"price"`
	City        string  `json:"city"`
	Photos      string  `json:"photos"`
}

// enviar mensajes a una cola RabbitMQ y notifica eventos de modificacion o creacion de un hotel
func (qc *QueueClient) SendMessage(id string, action string, message string /*hotelDetails HotelMessage*/) e.ApiError {

	channel, err := qc.Connection.Channel()

	err = channel.ExchangeDeclare(
		"hotels", // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return e.NewBadRequestApiError("Failed to declare an exchange")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	/*
		// Convierte los detalles del hotel a JSON
		body, jsonErr := json.Marshal(hotelDetails)
		if jsonErr != nil {
			return e.NewBadRequestApiError("Failed to encode hotel details as JSON")
		}
	*/

	body := message

	// Publica el mensaje en la cola
	err = channel.PublishWithContext(ctx,
		"hotels",
		fmt.Sprintf("%s.%s", id, action),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return e.NewBadRequestApiError("Failed to publish a message")
	}
	log.Printf("[x] Sent %s\n", body)
	return nil
}
