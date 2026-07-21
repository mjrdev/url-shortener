package config

import (
	"crypto/tls"
	"fmt"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	rdb     *redis.Client
	rdbOnce sync.Once
)

func Rdb() *redis.Client {
	rdbOnce.Do(func() {
		redisUrl := os.Getenv("REDIS_URL")
		redisPort := os.Getenv("REDIS_PORT")
		redisPassword := os.Getenv("REDIS_PASSWORD")

		options := &redis.Options{
			Addr:     fmt.Sprintf("%s:%s", redisUrl, redisPort),
			Password: redisPassword,
			DB:       0,
		}

		if os.Getenv("REDIS_TLS") == "true" {
			options.TLSConfig = &tls.Config{
				MinVersion: tls.VersionTLS12,
			}
		}

		rdb = redis.NewClient(options)
		fmt.Println("[CACHE] Connection sucess")
	})

	return rdb
}
