package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hotels/config"
	"hotels/dto"
	client "hotels/services/repositories"
	e "hotels/utils/errors"
	"net/http"
	"search/dto"

	json "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

type HotelService interface {
	GetHotelById(id string) (dto.HotelDto, e.ApiError)
	GetHotels() (dto.HotelDto, e.ApiError)
	InsertHotel(hotel dto.HotelDto) (dto.HotelDto, e.ApiError)
	QueueHotels(hotels dto.HotelsDto) e.ApiError
	DeleteHotel(id string) e.ApiError
	UpdateHotelById(hotel dto.HotelDto) (dto.HotelDto, e.ApiError)
}

type HotelServiceImpl struct {
	hotel *client.HotelClient
	queue *client.QueueClient
}

func NewHotelServiceImpl(
	hotel *client.HotelClient,
	queue *client.QueueClient,
) *HotelServiceImpl {
	return &HotelServiceImpl{
		hotel: hotel,
		queue: queue,
	}
}

func (s *HotelServiceImpl) GetHotelById(id string) (dto.ItemResponseDto, e.ApiError) {

	var hotelDto dto.HotelDto
	var hotelResponseDto dto.HotelResponseDto

	hotelDto, err := s.hotel.GetHotelById(id)

	if err != nil {
		log.Debug("Error getting hotel from mongo")
		return hotelResponseDto, err
	}

	if hotelDto.HotelId == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	log.Debug("mongo")
	return hotelDto

}

func (s *HotelServiceImpl) GetHotels() (dto.HotelsDto, e.ApiError) {

	var hotelsDto dto.HotelsDto
	var hotelsResponseDto dto.hotelsResponseDto

	hotelsDto, err := s.hotel.GetAllHotels()
	if err != nil {
		log.Debug("Error getting all hotels from mongo")
		return hotelsResponseDto, err
	}

	return hotelsDto, nil
}

// Inserta hoteles el la DB y notifica la cola RabbitMQ
func (s *HotelServiceImpl) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotelInsertDto dto.HotelDto

	hotelInsertDto, err := s.hotel.InsertHotel(hotelDto)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("error inserting hotel")
	}

	if hotelInsertDto.HotelId == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("error in insert")
	}

	hotelDto.HotelId = hotelInsertDto.HotelId

	//Inserta el ID del hotel en la cola RabbitMQ
	err = s.queue.QueueHotels(hotelDto.HotelId)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("Error notifying hotel creation to RabbitMQ")
	}

	hotelDto, err = s.queue.InsertHotel(hotelDto)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("Error inserting in queue")
	}
	return hotelDto, nil
}

// Actualiza la información del hotel en la base de datos,
// notifica la actualización a la cola de RabbitMQ
func (s *HotelServiceImpl) UpdateHotelById(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var updatedHotelDto dto.HotelDto

	// Actualiza el hotel en la base de datos
	updatedHotelDto, err := s.hotel.UpdateHotel(hotelDto)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("error updating hotel")
	}

	// Verifica si la actualización fue exitosa
	if updatedHotelDto.HotelId == "" {
		return hotelDto, e.NewBadRequestApiError("error in update")
	}

	// Inserta el ID del hotel actualizado en la cola RabbitMQ
	err = s.queue.QueueHotels(updatedHotelDto.HotelId)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("Error notifying hotel update to RabbitMQ")
	}

	return updatedHotelDto, nil
}

// Rastrea el procesamiento de elementos en una cola y notifica al
// usuario una vez que se han procesado todos los elementos
func CheckQueue(processed chan string, total int, userid int) {
	var complete int
	var errors int
	for loop := true; loop; {
		select {
		case data := <-processed:
			if data == "error" {
				errors++
			} else {
				complete++
			}
			if errors+complete == total {
				loop = false
			}
		default:
			log.Debugf("waiting for %d more messages", total-complete-errors)
		}
	}
	var body []byte
	var message dto.MessageDto
	message.UserId = userid
	message.System = true
	message.Body = fmt.Sprintf("Processed items total = %d, Completed: %d, Errors: %d", complete+errors, complete, errors)
	body, err := json.Marshal(&message)

	if err != nil {
		panic(e.NewInternalServerApiError("Error marshaling in sending message", err))
	}
	_, err = http.Post(fmt.Sprintf("http://%s:%d/%s", config.MESSAGESHOST, config.MESSAGESPORT, config.MESSAGESENDPOINT), "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(e.NewInternalServerApiError("Error sending message to message service", err))

	}
}

// Notifica a una cola de mensajes (RabbitMQ) cuando se crea o modifica un hotel
func (s *HotelServiceImpl) QueueHotels(hotelsDto dto.HotelsDto) e.ApiError {
	total := len(hotelsDto)
	processed := make(chan string, total)

	for i := range hotelsDto {
		var hotel dto.HotelDto
		hotel = hotelsDto[i]
		go func() {

			// Insertar el hotel en MongoDB
			hotel, err := s.hotel.InsertHotel(hotel)
			if err != nil {
				processed <- "error"
				log.Debug(err)
				return
			}

			// Notificar a RabbitMQ cuando se crea o modifica un hotel
			var action string
			if hotel.HotelId == "" {
				action = "create"
			} else {
				action = "update"
			}
			err = s.queue.SendMessage(hotel.HotelId, action, hotel.HotelId)
			if err != nil {
				processed <- "error"
				log.Debug(err)
				return
			}

			processed <- "complete"
		}()
	}

	go CheckQueue(processed, total, hotelsDto[0].UsuarioId)
	return nil
}

// Elimina un hotel de la base de datos y notificar la eliminación a través de una cola de mensajes
func (s *HotelServiceImpl) DeleteHotelById(id string) e.ApiError {

	err := s.hotel.DeleteHotel(id)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.queue.SendMessage(id, "delete", fmt.Sprintf("%s.delete", id))
	log.Error(err)

	return nil
}
