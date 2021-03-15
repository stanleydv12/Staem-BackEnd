package helper

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var redisInstance redis.UniversalClient

func getCache() redis.UniversalClient {
	if redisInstance == nil {
		option, err := redis.ParseURL("redis://:fc92f56284084f428e411397410e64c2@us1-merry-boar-32294.upstash.io:32294")
		if err != nil {
			log.Fatal(err)
		}

		redisInstance = redis.NewClient(option)
	}
	return redisInstance
}

func UseCache() redis.UniversalClient {
	return getCache()
}

func UseCached(ctx context.Context, key string, cacheFunction func() (string, error)) (string, error) {
	cache := getCache()

	if err := cache.Ping(ctx).Err(); err != nil {
		return "", err
	}

	cached, err := cache.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}

	if cached != "" {
		log.Printf("Redis: %v found with value %v", key, cached)
		return cached, nil
	}

	log.Printf("Redis: %v doesn't exist (%v), creating a new one...", key, cached)

	value, err := cacheFunction()
	if err != nil {
		return "", err
	}

	return value, cache.Set(ctx, key, cached, 30*time.Second).Err()
}
