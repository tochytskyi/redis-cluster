package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

/**

https://en.wikipedia.org/wiki/Cache_stampede

function x-fetch(key, ttl, beta=1) {
    value, delta, expiry ← cache_read(key)
    if (!value || (time() - delta * beta * log(rand(0,1)) ≥ expiry) {
        start ← time()
        value ← recompute_value()
        delta ← time() – start
        cache_write(key, (value, delta), ttl)
    }
    return value
}
*/

type CacheItem struct {
	Value string `json:"value"`
	Delta int64  `json:"delta"`
}

var beta int = 1

func checkStampedeFactor(valueObject CacheItem, expirySeconds int) bool {
	return float64(valueObject.Delta)*float64(beta)*math.Log(rand.Float64()) >= float64(expirySeconds)
}

func GetOrRefresh(key string, ttl int) (string, error) {
	log.Println("Looking for key: " + key)

	var expirySeconds int = 0

	value, err := GetInstance().Get(key).Result()
	var valueObject CacheItem

	if err == nil {
		log.Println("Data in cache", value)

		ttlFromCache, _ := GetInstance().TTL(key).Result()
		expirySeconds := ttlFromCache.Seconds()

		if expirySeconds < 0 {
			if -1 == expirySeconds {
				log.Println("The key will not expire")
				return "", fmt.Errorf("The key will not expire")
			} else if -2 == expirySeconds {
				log.Println("The key does not exist")
				return "", fmt.Errorf("The key does not exist")
			} else {
				log.Printf("Unexpected error %d", expirySeconds)
				return "", fmt.Errorf("Unexpected error %d", expirySeconds)
			}
		} else {
			log.Println("Current TTL", expirySeconds)
		}

		json.Unmarshal([]byte(value), &valueObject)

	} else {
		log.Println("No data in cache", err)
	}

	if (valueObject == CacheItem{} || checkStampedeFactor(valueObject, expirySeconds)) {
		recomputedValue := "new value"

		recomputedValueString, err := json.Marshal(CacheItem{
			Value: recomputedValue,
			Delta: rand.Int63n(10) + 1, //simulate cache recalculating time delay
		})

		if err != nil {
			log.Println("Failed to marshal new value", err)
			return "", err
		}

		GetInstance().Set(key, recomputedValueString, time.Duration(ttl)).Result()

		log.Println("New value cached", string(recomputedValueString))
	}

	return value, nil
}
