package services

import (
	"fmt"
	"hotels/dto"

	client "hotels/services/repositories"

	e "hotels/utils/errors"

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

//get hotelid de antes por si acaso
/*
func (s *HotelServiceImpl) GetHotelById(id string) (dto.HotelResponseDto, e.ApiError) {

	var hotelDto dto.HotelDto
	var hotelResponseDto dto.HotelResponseDto

	hotelDto, err := s.hotel.GetHotelById(id)

	if err != nil {
		log.Debug("Error getting hotel from mongo")
		return hotelResponseDto, err
	}

	if hotelDto.HotelId == "000000000000000000000000" {
		return hotelResponseDto, e.NewBadRequestApiError("hotel not found")
	}

	log.Debug("mongo")

	return hotelResponseDto, nil

}*/

func (s *HotelServiceImpl) GetHotelById(id string) (dto.HotelDto, e.ApiError) {

	var hotelDto dto.HotelDto
	//var hotelResponseDto dto.HotelResponseDto

	hotelDto, err := s.hotel.GetHotelById(id)

	if err != nil {
		log.Debug("Error getting hotel from mongo")
		return hotelDto, err
	}

	if hotelDto.HotelId == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	log.Debug("mongo")

	return hotelDto, nil

}

func (s *HotelServiceImpl) GetHotels() (dto.HotelsDto, e.ApiError) {

	var hotelsDto dto.HotelsDto

	hotelsDto, err := s.hotel.GetAllHotels()
	if err != nil {
		log.Debug("Error getting all hotels from mongo")
		return hotelsDto, err
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
	err = s.queue.SendMessage(hotelDto.HotelId, "create", "") //err = s.queue.QueueHotels(hotelDto.HotelId)

	if err != nil {
		return hotelDto, e.NewBadRequestApiError("Error notifying hotel creation to RabbitMQ")
	}

	err = s.queue.SendMessage(hotelDto.HotelId, "update", "") //hotelDto, err = s.queue.InsertHotel(hotelDto)

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
	updatedHotelDto, err := s.hotel.UpdateHotelById(hotelDto)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("error updating hotel")
	}

	// Verifica si la actualización fue exitosa
	if updatedHotelDto.HotelId == "" {
		return hotelDto, e.NewBadRequestApiError("error in update")
	}

	// Inserta el ID del hotel actualizado en la cola RabbitMQ
	err = s.queue.SendMessage(updatedHotelDto.HotelId, "update", "") //err = s.queue.QueueHotels(updatedHotelDto.HotelId)

	if err != nil {
		return hotelDto, e.NewBadRequestApiError("Error notifying hotel update to RabbitMQ")
	}

	return updatedHotelDto, nil
}

func (s *HotelServiceImpl) QueueHotels(hotelsDto dto.HotelsDto) e.ApiError {
	// Crea un canal para mantener el seguimiento del progreso de procesamiento.
	processed := make(chan string, len(hotelsDto))

	for i := range hotelsDto {
		hotel := hotelsDto[i]
		go func() {

			// Variable para almacenar la acción (create o update).
			var action string

			// Insertar el hotel en MongoDB
			hotel, err := s.hotel.InsertHotel(hotel)
			if err != nil {
				processed <- "error"
				log.Debug(err)
				return
			}

			// Define la acción en función de si se creó o actualizó el hotel
			if hotel.HotelId == "" {
				action = "create"
			} else {
				action = "update"
			}

			// Notificar a RabbitMQ cuando se crea o modifica un hotel
			err = s.queue.SendMessage(hotel.HotelId, action, "")
			if err != nil {
				processed <- "error"
				log.Debug(err)
				return
			}

			processed <- "complete"
		}()
	}

	// Notifica cuando se completa la operación.
	go func() {
		// Verifica si hubo errores en el procesamiento antes de notificar.
		completeCount := 0
		errorCount := 0
		for i := 0; i < len(hotelsDto); i++ {
			result := <-processed
			if result == "error" {
				errorCount++
			} else if result == "complete" {
				completeCount++
			}
		}

		// Determina si hubo errores de procesamiento
		action := "processed"
		if errorCount > 0 {
			action = "processed with errors"
		}

		// Llama a SendMessage para notificar la operación
		err := s.queue.SendMessage("all-hotels", action, fmt.Sprintf("Processed %d hotels with %d errors", completeCount, errorCount))
		if err != nil {
			log.Debug(err)
		}
	}()

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
