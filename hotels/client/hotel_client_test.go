package client_test

import (
	"errors"
	"hotels/client"
	"hotels/db"
	"hotels/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsertHotel(t *testing.T) {

	a := assert.New(t)

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		db.HotelsCollection = mt.Coll

		hotel := client.HotelClient.InsertHotel(model.Hotel{
			Name:        "Hotel Test",
			Description: "Hotel test description",
			Amenities:   "Test Amenities",
			Stars:       1,
			Rooms:       10,
			Price:       10.2,
			City:        "Test City",
		})

		a.NotNil(hotel.HotelId)
	})
}

func TestGetHotelById(t *testing.T) {

	a := assert.New(t)

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {

		db.HotelsCollection = mt.Coll

		expectedHotel := model.Hotel{
			HotelId:     primitive.NewObjectID(),
			Name:        "Hotel Test",
			Description: "Hotel test description",
			Amenities:   "Test Amenities",
			Stars:       1,
			Rooms:       10,
			Price:       10.2,
			City:        "Test City",
		}

		cursor := mtest.CreateCursorResponse(1, "hotel.1", mtest.FirstBatch, bson.D{
			{"_id", expectedHotel.HotelId},
			{"name", expectedHotel.Name},
			{"description", expectedHotel.Description},
			{"amenities", expectedHotel.Amenities},
			{"stars", expectedHotel.Stars},
			{"rooms", expectedHotel.Rooms},
			{"price", expectedHotel.Price},
			{"city", expectedHotel.City},
		})

		mt.AddMockResponses(cursor)

		hotel, err := client.HotelClient.GetHotelById(expectedHotel.HotelId.Hex())

		if err != nil {
			err = errors.New("error retrieving hotel by ID: " + err.Error())
			return
		}

		a.Equal(expectedHotel, hotel)
	})

	mt.Run("failure", func(mt *mtest.T) {

		db.HotelsCollection = mt.Coll

		hotel, err := client.HotelClient.GetHotelById(primitive.NewObjectID().Hex())

		if err != nil {
			err = errors.New("error retrieving hotel by ID: " + err.Error())
			return
		}

		var emptyModel model.Hotel

		a.Equal(emptyModel, hotel)
	})
}

func TestGetAllHotels(t *testing.T) {

	a := assert.New(t)

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {

		db.HotelsCollection = mt.Coll

		expectedHotel1 := model.Hotel{
			HotelId:     primitive.NewObjectID(),
			Name:        "Hotel Test 1",
			Description: "Hotel test 1 description",
			Amenities:   "Test Amenities 1",
			Stars:       1,
			Rooms:       1,
			Price:       1.1,
			City:        "Test City 1",
		}

		expectedHotel2 := model.Hotel{
			HotelId:     primitive.NewObjectID(),
			Name:        "Hotel Test 2",
			Description: "Hotel test 2 description",
			Amenities:   "Test Amenities 2",
			Stars:       2,
			Rooms:       2,
			Price:       2.2,
			City:        "Test City 2",
		}

		cursor1 := mtest.CreateCursorResponse(1, "hotel.1", mtest.FirstBatch, bson.D{
			{"_id", expectedHotel1.HotelId},
			{"name", expectedHotel1.Name},
			{"description", expectedHotel1.Description},
			{"amenities", expectedHotel1.Amenities},
			{"stars", expectedHotel1.Stars},
			{"rooms", expectedHotel1.Rooms},
			{"price", expectedHotel1.Price},
			{"city", expectedHotel1.City},
		})

		cursor2 := mtest.CreateCursorResponse(1, "hotel.1", mtest.NextBatch, bson.D{
			{"_id", expectedHotel2.HotelId},
			{"name", expectedHotel2.Name},
			{"description", expectedHotel2.Description},
			{"amenities", expectedHotel2.Amenities},
			{"stars", expectedHotel2.Stars},
			{"rooms", expectedHotel2.Rooms},
			{"price", expectedHotel2.Price},
			{"city", expectedHotel2.City},
		})

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)

		mt.AddMockResponses(cursor1, cursor2, killCursors)

		hotels := client.HotelClient.GetAllHotels()

		var expectedHotels model.Hotels

		expectedHotels = append(expectedHotels, expectedHotel1)
		expectedHotels = append(expectedHotels, expectedHotel2)

		a.Equal(expectedHotels, hotels)

	})
}

func TestDeleteHotelById(t *testing.T) {

	a := assert.New(t)

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		db.HotelsCollection = mt.Coll

		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})

		err := client.HotelClient.DeleteHotelById(primitive.NewObjectID().Hex())
		a.Nil(err)
	})

	mt.Run("failure", func(mt *mtest.T) {
		db.HotelsCollection = mt.Coll

		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})

		err := client.HotelClient.DeleteHotelById(primitive.NewObjectID().Hex())
		a.NotNil(err)

		expectedResponse := "hotel not found"
		a.Equal(expectedResponse, err.Error())
	})

	mt.Run("error", func(mt *mtest.T) {
		db.HotelsCollection = mt.Coll

		mt.AddMockResponses(bson.D{{"ok", 0}})

		err := client.HotelClient.DeleteHotelById(primitive.NewObjectID().Hex())
		a.NotNil(err)

		expectedResponse := "command failed"
		a.Equal(expectedResponse, err.Error())
	})
}

func TestUpdateHotelById(t *testing.T) {

	a := assert.New(t)

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {

		db.HotelsCollection = mt.Coll

		updatedHotel := model.Hotel{
			HotelId:     primitive.NewObjectID(),
			Name:        "Hotel Test 1",
			Description: "Hotel test 1 description",
			Amenities:   "Test Amenities 1",
			Stars:       1,
			Rooms:       1,
			Price:       1.1,
			City:        "Test City 1",
		}

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"n", 1},
		})

		hotelResponse := client.HotelClient.UpdateHotelById(updatedHotel.HotelId.Hex(), updatedHotel)

		a.Equal(updatedHotel, hotelResponse)

	})

	mt.Run("failure", func(mt *mtest.T) {

		db.HotelsCollection = mt.Coll

		updatedHotel := model.Hotel{
			HotelId:     primitive.NewObjectID(),
			Name:        "Hotel Test 1",
			Description: "Hotel test 1 description",
			Amenities:   "Test Amenities 1",
			Stars:       1,
			Rooms:       1,
			Price:       1.1,
			City:        "Test City 1",
		}

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"n", 0},
		})

		hotelResponse := client.HotelClient.UpdateHotelById(updatedHotel.HotelId.Hex(), updatedHotel)

		emptyModel := model.Hotel{}

		a.Equal(emptyModel, hotelResponse)

	})
}
