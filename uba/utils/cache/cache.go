package utils

import (
	"github.com/bradfitz/gomemcache/memcache"

	log "github.com/sirupsen/logrus"
)

type CacheInterface interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl int32) error
}

type Cache struct {
	client *memcache.Client
}

func InitCache() *Cache {
	client := memcache.New("memcached:11211")
	log.Info("Initializing cache")
	return NewCache(client)
}

func NewCache(client *memcache.Client) *Cache {
	return &Cache{
		client: client,
	}
}

func (c *Cache) Get(key string) ([]byte, error) {
	item, err := c.client.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			log.Printf("Cache miss for key: %s", key)
		} else {
			log.Printf("Error getting key %s from cache: %v", key, err)
		}
		return nil, err
	}
	log.Printf("Cache hit for key: %s", key)
	return item.Value, nil
}

func (c *Cache) Set(key string, value []byte, ttl int32) error {
	err := c.client.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: ttl,
	})
	if err != nil {
		log.Printf("Error setting key %s to cache: %v", key, err)
		return err
	}
	log.Printf("Key %s set in cache with TTL: %d", key, ttl)
	return nil
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
