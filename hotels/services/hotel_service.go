package services

import (
	"hotels/dto"
	e "hotels/utils/errors"
)

type HotelService interface {
	GetHotelById(id string) (dto.HotelDto, e.ApiError)
	InsertItem(hotel dto.HotelDto) (dto.HotelDto, e.ApiError)
	QueueItems(hotels dto.HotelsDto) e.ApiError
	DeleteUserItems(id int) e.ApiError
	DeleteItem(id string) e.ApiError
}

type ItemServiceImpl struct {
	hotel      *client.HotelClient
	queue     *client.QueueClient
}

func NewItemServiceImpl(
	item *client.ItemClient,
	queue *client.QueueClient,
) *ItemServiceImpl {
	return &ItemServiceImpl{
		item:      item,
		queue:     queue,
	}
}


func (s *ItemServiceImpl) GetItemById(id string) (dto.ItemResponseDto, e.ApiError) {

	var itemDto dto.ItemDto
	var itemResponseDto dto.ItemResponseDto

	itemDto, err := s.memcached.GetItemById(id)
	if err != nil {
		log.Debug("Error getting item from memcached")
		itemDto, err2 := s.item.GetItemById(id)
		if err2 != nil {
			log.Debug("Error getting item from mongo")
			return itemResponseDto, err2
		}
		if itemDto.ItemId == "000000000000000000000000" {
			return itemResponseDto, e.NewBadRequestApiError("item not found")
		}
		_, err3 := s.memcached.InsertItem(itemDto)
		if err3 != nil {
			log.Debug("Error inserting in memcached")
		}
		log.Debug("mongo")
		return s.GetUserById(itemDto.UsuarioId, itemDto)
	} else {
		log.Debug("memcached")
		return s.GetUserById(itemDto.UsuarioId, itemDto)
	}
}

func (s *ItemServiceImpl) GetItemsByUserId(id int) (dto.ItemsResponseDto, e.ApiError) {

	var itemsDto dto.ItemsDto
	var itemsResponseDto dto.ItemsResponseDto
	itemsDto, err := s.item.GetItemsByUserId(id)
	if err != nil {
		log.Debug("Error getting items from mongo")
		return itemsResponseDto, err
	}

	for i := range itemsDto {
		item, err := s.GetUserById(itemsDto[i].UsuarioId, itemsDto[i])
		if err != nil {
			return itemsResponseDto, e.NewBadRequestApiError("error getting user for item")
		}
		itemsResponseDto = append(itemsResponseDto, item)
	}

	return itemsResponseDto, nil

}

func (s *ItemServiceImpl) InsertItem(itemDto dto.ItemDto) (dto.ItemDto, e.ApiError) {

	var insertItem dto.ItemDto

	insertItem, err := s.item.InsertItem(itemDto)
	if err != nil {
		return itemDto, e.NewBadRequestApiError("error inserting item")
	}

	if insertItem.ItemId == "000000000000000000000000" {
		return itemDto, e.NewBadRequestApiError("error in insert")
	}
	itemDto.ItemId = insertItem.ItemId

	itemDto, err = s.memcached.InsertItem(itemDto)
	if err != nil {
		return itemDto, e.NewBadRequestApiError("Error inserting in memcached")
	}
	return itemDto, nil
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

func (s *ItemServiceImpl) QueueItems(itemsDto dto.ItemsDto) e.ApiError {
	total := len(itemsDto)
	processed := make(chan string, total)
	for i := range itemsDto {
		var item dto.ItemDto
		item = itemsDto[i]
		go func() {
			url := item.UrlImg
			item, err := s.item.InsertItem(item)
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

func (s *ItemServiceImpl) DeleteUserItems(id int) e.ApiError {
	items, err := s.GetItemsByUserId(id)
	if err != nil {
		log.Error(err)
		return err
	}
	for i := range items {
		var item dto.ItemResponseDto
		item = items[i]
		go func() {
			err := s.item.DeleteItem(item.ItemId)
			if err != nil {
				log.Error(err)
			}
			err = s.queue.SendMessage(item.ItemId, "delete", fmt.Sprintf("%s.delete", item.ItemId))
			log.Error(err)
		}()
	}
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