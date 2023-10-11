package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hotels/config"
	"hotels/dto"
	client "hotels/services/repositories"
	e "hotels/utils/errors"
	"io"
	"net/http"
	"os"
	"strings"

	json "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

type HotelService interface {
	GetHotelById(id string) (dto.HotelDto, e.ApiError)
	InsertHotel(hotel dto.HotelDto) (dto.HotelDto, e.ApiError)
	QueueHotels(hotels dto.HotelsDto) e.ApiError
	DeleteHotel(id string) e.ApiError
}

type HotelServiceImpl struct {
	hotel *client.HotelClient
	queue *client.QueueClient
}

func NewHotelServiceImpl(
	item *client.HotelClient,
	queue *client.QueueClient,
) *HotelServiceImpl {
	return &HotelServiceImpl{
		hotel: hotel,
		queue: queue,
	}
}

func (s *HotelServiceImpl) GetHotelById(id string) (dto.ItemResponseDto, e.ApiError) {

	var hotelDto dto.HotelDto
	var itemResponseDto dto.ItemResponseDto

	hotelDto, err := s.hotel.GetHotelById(id)

	if err != nil {
		log.Debug("Error getting hotel from mongo")
		return itemResponseDto, err
	}

	if hotelDto.HotelId == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	log.Debug("mongo")
	return hotelDto

}

func (s *HotelServiceImpl) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotelInsertDto dto.HotelDto

	hotelInsertDto, err := s.hotel.InsertHotel(hotelDto)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("error inserting hotel")
	}

	if hotelInsertDto.ItemId == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("error in insert")
	}

	hotelDto.ItemId = hotelInsertDto.ItemId

	//HERE: INSERT TO QUEUE !! CHECK

	hotelDto, err = s.queue.InsertItem(hotelDto)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError("Error inserting in queue")
	}
	return hotelDto, nil
}

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

func DownloadImage(url string, name string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	path := strings.Join([]string{"/exports/images", name}, "/")
	file, _ := os.Create(path)
	defer file.Close()
	_, _ = io.Copy(file, resp.Body)
}

func (s *HotelServiceImpl) QueueHotels(itemsDto dto.ItemsDto) e.ApiError {
	total := len(itemsDto)
	processed := make(chan string, total)
	for i := range itemsDto {
		var item dto.ItemDto
		item = itemsDto[i]
		go func() {
			url := item.UrlImg
			item, err := s.hotel.InsertItem(item)
			go DownloadImage(url, item.UrlImg)
			log.Debug(url)
			log.Debug(item.UrlImg)
			if err != nil {
				processed <- "error"
				log.Debug(err)
			}
			processed <- "complete"
			err = s.queue.SendMessage(item.ItemId, "create", item.ItemId)
			log.Debug(err)
		}()
	}

	go CheckQueue(processed, total, itemsDto[0].UsuarioId)
	return nil
}

func (s *ItemServiceImpl) DeleteItemById(id string) e.ApiError {

	err := s.item.DeleteItem(id)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.memcached.DeleteItem(id)
	if err != nil {
		log.Error("Error deleting from cache", err)
	}
	err = s.queue.SendMessage(id, "delete", fmt.Sprintf("%s.delete", id))
	log.Error(err)

	return nil
}
