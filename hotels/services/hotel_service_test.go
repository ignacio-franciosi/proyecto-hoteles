package services_test

import (
	"errors"
	"hotels/client"
	"hotels/dto"
	"hotels/model"
	"hotels/queue"
	"hotels/services"
	"net/http"
	"testing"

	"github.com/NeowayLabs/wabbit"
	"github.com/NeowayLabs/wabbit/amqptest"
	"github.com/NeowayLabs/wabbit/amqptest/server"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type testHotel struct{}
type testQueue struct{}
type mockHTTPClient struct{}

var channel wabbit.Channel
var mockQueue wabbit.Queue

func init() {
	client.HotelClient = testHotel{}
	queue.QueueProducer = testQueue{}
	queue.QueueProducer.InitQueue()
	//services.HotelService = &hotelService{HTTPClient: &mockHTTPClient{}}
}

func (m *mockHTTPClient) Do(_ *http.Request) (*http.Response, error) {

	return &http.Response{
		StatusCode: http.StatusOK,
	}, nil
}

func (t testHotel) InsertHotel(hotel model.Hotel) model.Hotel {

	if hotel.Name != "" {
		objId, _ := primitive.ObjectIDFromHex("654cf68d807298d99186019f")
		hotel.HotelId = objId
	}

	return hotel
}

func (t testHotel) GetHotelById(id string) (model.Hotel, error) {
	if id == "000000000000000000000000" {
		return model.Hotel{}, errors.New("invalid hotel ID provided")
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Hotel{}, errors.New("error converting ID to ObjectID: " + err.Error())
	}

	return model.Hotel{
		HotelId: objId,
	}, nil
}

func (t testHotel) GetAllHotels() model.Hotels {
	return model.Hotels{}
}

func (t testHotel) DeleteHotelById(id string) error {

	if id == "000000000000000000000000" {
		return errors.New("hotel not found")
	}

	return nil
}

func (t testHotel) UpdateHotelById(id string, hotel model.Hotel) model.Hotel {
	if id == "000000000000000000000000" {
		return model.Hotel{}
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Hotel{}
	}

	hotel.HotelId = objId
	return hotel
}

func (t testQueue) InitQueue() {

	fakeServer := server.NewServer("amqp://localhost:5672/%2f")
	err := fakeServer.Start()

	if err != nil {
		log.Info("Failed to start server")
		log.Fatal(err)
	}

	connection, err := amqptest.Dial("amqp://localhost:5672/%2f")

	if err != nil {
		log.Info("Failed to connect to RabbitMQ")
		log.Fatal(err)
	} else {
		log.Info("RabbitMQ connection established")
	}

	channel, err = connection.Channel()

	if err != nil {
		log.Info("Failed to open channel")
		log.Fatal(err)
	}

	defer channel.Close()

	mockQueue, err = channel.QueueDeclare("hotel", wabbit.Option{
		"durable":    true,
		"autoDelete": false,
		"exclusive":  false,
		"noWait":     false,
	})

	if err != nil {
		log.Info("Failed to declare a queue")
		log.Fatal(err)
	} else {
		log.Info("Queue declared")
	}
}

func (t testQueue) Publish(body []byte) error {

	err := channel.Publish(
		"",
		mockQueue.Name(),
		body,
		wabbit.Option{"contentType": "application/json"})

	if err != nil {
		log.Debug("Error while publishing message", err)
		return err
	}

	return nil
}

func TestInsertHotel(t *testing.T) {

	a := assert.New(t)

	hotelDto := dto.HotelDto{
		Name:        "Hotel Test",
		Description: "Hotel test description",
		Amenities:   "Test Amenities",
		Stars:       1,
		Rooms:       10,
		Price:       10.2,
		City:        "Test City",
	}

	hotelResponse, err := services.HotelService.InsertHotel(hotelDto)

	hotelDto.HotelId = "654cf68d807298d99186019f"

	a.Nil(err)
	a.Equal(hotelDto, hotelResponse)

}

func TestGetHotelById(t *testing.T) {

	a := assert.New(t)

	id := "654cf68d807298d99186019f"

	hotelResponse, err := services.HotelService.GetHotelById(id)

	a.Nil(err)
	a.Equal(id, hotelResponse.HotelId)
}

func TestGetAllHotels(t *testing.T) {

	a := assert.New(t)

	hotelsResponse, err := services.HotelService.GetAllHotels()

	a.Nil(err)

	var emptyDto dto.HotelsDto
	a.Equal(emptyDto, hotelsResponse)
}

func TestDeleteHotel(t *testing.T) {

	a := assert.New(t)

	id := "654cf68d807298d99186019f"
	_, err := services.HotelService.DeleteHotel(id)

	a.Nil(err)
}

func TestUpdateHotel(t *testing.T) {
	a := assert.New(t)

	id := "654cf68d807298d99186019f"

	hotelDto := dto.HotelDto{HotelId: id}

	hotelResponse, err := services.HotelService.UpdateHotel(hotelDto)

	a.Nil(err)
	a.Equal(hotelDto, hotelResponse)
}
