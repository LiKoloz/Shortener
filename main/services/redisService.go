package services

import (
	"Shortener/models"
	"context"
	"time"

	"github.com/wb-go/wbf/redis"
	"github.com/wb-go/wbf/retry"
)

var client *redis.Client = nil

// Get получение из рредиса
func Get(shortUrl string) (string, error) {
	if client == nil {
		client = redis.New("redis:6379", "", 0)
	}
	strategy := retry.Strategy{Attempts: 3, Delay: 5 * time.Second, Backoff: 2}
	ctx := context.Background()
	val, err := client.GetWithRetry(ctx, strategy, shortUrl)
	if err != nil {
		return "", err
	}

	return val, nil
}

// Add добавление в редис
func Add(url models.Url) error {
	if client == nil {
		client = redis.New("redis:6379", "", 0)
	}
	strategy := retry.Strategy{Attempts: 3, Delay: 5 * time.Second, Backoff: 2}
	key := url.Short
	value := url.Long
	expiration := time.Hour
	ctx := context.Background()
	if err := client.SetWithExpirationAndRetry(ctx, strategy, key, value, expiration); err != nil {
		return err
	}
	return nil
}
