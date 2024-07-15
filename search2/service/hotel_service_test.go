package service_test

import (
	"errors"
	"search2/client"
	"search2/dto"
	"search2/service"
	"testing"

	solr "github.com/rtt/Go-Solr"
	"github.com/stretchr/testify/assert"
)

type testHotel struct{}

func init() {
	client.SolrHotelClient = testHotel{}
}

func (t testHotel) UpdateHotel(document map[string]interface{}) error {

	var id string

	add, exists := document["add"].([]interface{})

	if exists {
		id = add[0].(map[string]interface{})["id"].(string)
	} else {
		id = document["delete"].([]interface{})[0].(map[string]interface{})["id"].(string)
	}

	if id == "" {
		return errors.New("error updating hotel")
	}

	return nil
}

func (t testHotel) GetHotels() (*solr.DocumentCollection, error) {

	doc := solr.DocumentCollection{}

	doc.Collection = make([]solr.Document, 0)

	newDoc1 := solr.Document{
		Fields: map[string]interface{}{
			"id":          "1",
			"hotel_id":    []string{"654cf68d807298d99186019f"},
			"name":        []string{"Hotel Test 1"},
			"rooms":       []int{10},
			"description": []string{"Hotel Test Description 1"},
			"city":        []string{"Test City 1"},
			"stars":       []int{123},
			"price":       []float64{4.5},
			"amenities":   []string{"Test Amenities 1"},
		},
	}

	newDoc2 := solr.Document{
		Fields: map[string]interface{}{
			"id":          "2",
			"hotel_id":    []string{"654cf68d807298d99186019g"},
			"name":        []string{"Hotel Test 2"},
			"rooms":       []int{11},
			"description": []string{"Hotel Test Description 2"},
			"city":        []string{"Test City 2"},
			"stars":       []int{124},
			"price":       []float64{4.6},
			"amenities":   []string{"Test Amenities 2"},
		},
	}

	doc.Collection = append(doc.Collection, newDoc1)
	doc.Collection = append(doc.Collection, newDoc2)

	return &doc, nil
}

func (t testHotel) GetHotelById(id string) (*solr.DocumentCollection, error) {

	if id == "000000000000000000000000" {
		return &solr.DocumentCollection{}, errors.New("hotel not found")
	}

	doc := solr.DocumentCollection{}

	doc.Collection = make([]solr.Document, 0)

	newDoc := solr.Document{
		Fields: map[string]interface{}{
			"id":          "1",
			"hotel_id":    []string{"654cf68d807298d99186019f"},
			"name":        []string{"Hotel Test 1"},
			"rooms":       []int{10},
			"description": []string{"Hotel Test Description 1"},
			"city":        []string{"Test City 1"},
			"stars":       []int{123},
			"price":       []float64{4.5},
			"amenities":   []string{"Test Amenities 1"},
		},
	}

	doc.Collection = append(doc.Collection, newDoc)

	return &doc, nil
}

func (t testHotel) GetHotelsByCity(city string) (*solr.DocumentCollection, error) {

	if city == "" {
		return &solr.DocumentCollection{}, errors.New("city not found")
	}

	doc := solr.DocumentCollection{}
	doc.Collection = make([]solr.Document, 0)

	return &doc, nil
}

func TestInsertHotel(t *testing.T) {

	a := assert.New(t)

	hotelDto := dto.HotelDto{
		HotelId:     "654cf68d807298d99186019f",
		Name:        "Hotel Test",
		Rooms:       10,
		Description: "Hotel test Description",
		City:        "Test City",
		Stars:       123,
		Price:       4.5,
		Amenities:   "Test amenities",
	}

	err := service.HotelService.InsertUpdateHotel(hotelDto)
	a.Nil(err)

}

func TestGetHotelById(t *testing.T) {

	a := assert.New(t)

	id := "654cf68d807298d99186019f"

	hotelResponse, err := service.HotelService.GetHotelById(id)

	a.Nil(err)
	a.Equal(id, hotelResponse.HotelId)
}

func TestGetHotels(t *testing.T) {

	a := assert.New(t)

	hotelsResponse, err := service.HotelService.GetHotels()

	a.Nil(err)

	hotelsDto := dto.HotelsDto{
		dto.HotelDto{
			HotelId:     "654cf68d807298d99186019f",
			Name:        "Hotel Test 1",
			Rooms:       10,
			Description: "Hotel Test Description 1",
			City:        "Test City 1",
			Stars:       123,
			Price:       4.5,
			Amenities:   "Test Amenities 1",
		},

		dto.HotelDto{
			HotelId:     "654cf68d807298d99186019g",
			Name:        "Hotel Test 2",
			Rooms:       11,
			Description: "Hotel Test Description 2",
			City:        "Test City 2",
			Stars:       124,
			Price:       4.6,
			Amenities:   "Test Amenities 2",
		},
	}

	a.Equal(hotelsDto, hotelsResponse)
}

func TestGetHotelByCity(t *testing.T) {

	a := assert.New(t)

	city := "test"
	hotelsResponse, err := service.HotelService.GetHotelByCity(city)
	a.Nil(err)

	var emptyDto dto.HotelsDto
	a.Equal(emptyDto, hotelsResponse)
}

func TestDeleteHotelById(t *testing.T) {

	a := assert.New(t)

	id := "654cf68d807298d99186019f"

	err := service.HotelService.DeleteHotelById(id)

	a.Nil(err)
}
