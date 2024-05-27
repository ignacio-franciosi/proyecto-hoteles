package utils

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"

	"uba/dto"
	e "uba/utils/errors"

	json "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

var (
	cacheClient *memcache.Client
)

func Init_Cache() {
	cacheClient = memcache.New("cache")
	fmt.Println("Initialized cache", cacheClient)
	log.Info("Initialized cache")
}

func Set(key string, value []byte, ttlSeconds int) {

	//key := createCacheKey(id, startDate)
	//key := strconv.Itoa(id) + strconv.Itoa(startDate)
	if err := cacheClient.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(ttlSeconds), // Tiempo en segundos antes de que el elemento expire
	}); err != nil {
		fmt.Println("Error setting item in cache", err)
	}

}

func Get(key string) (dto.HotelsDto, e.ApiError) {
	fmt.Println("entro")
	response, err := cacheClient.Get(key)
	fmt.Println("obtuvo la key")
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return dto.HotelsDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", key))
		}
		errorMsg := fmt.Sprintf("Error getting item from cache: %s", key)
		fmt.Println(errorMsg)
		return dto.HotelsDto{}, e.NewInternalServerApiError(errorMsg, err)
	}
	var responseDto dto.HotelsDto
	if err := json.Unmarshal(response.Value, &responseDto); err != nil {
		return dto.HotelsDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", key), err)

	}
	return responseDto, nil
}

// func createCacheKey(id int, startDate int) string {
//     return fmt.Sprintf("reservation:%d:%d", id, startDate)
// }
// func main() {
//     Init_cache()
//     value := []byte("some data")
//     Set(1, 20231009, value)
//     result := Get(1, 20231009)
//     fmt.Printf("Result: %s\n", string(result))
// }

/*package utils

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	log "github.com/sirupsen/logrus"
)

var cache *memcache.Client

type Cache struct{}

type CacheInterface interface {
	Set(key string, value []byte)
	Get(key string) ([]byte, error)
}

func InitCache() {
	cache = memcache.New("memcached:11211")
	fmt.Println("Initialized cache", cache)
	log.Info("Initialized cache")
}

func (c *Cache) Set(key string, value []byte) {
	//aca le seteamos el ttl 10 seg
	err := cache.Set(&memcache.Item{Key: key, Value: value, Expiration: 10})

	if err != nil {
		log.Error("Error setting cache: ", err)
	}
}

func (c *Cache) Get(key string) ([]byte, error) {

	resp, err := cache.Get(key)

	if err != nil {
		log.Error("Error getting cache data: ", err)
		return []byte{}, err
	}

	log.Info("Retrieved data from cache")
	return resp.Value, nil
}
*/
