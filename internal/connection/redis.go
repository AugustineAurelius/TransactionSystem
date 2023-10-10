package connection

import (
	"TransactionSystem/internal/config"
	"github.com/redis/go-redis/v9"
)

var RedisVault *redis.Client

func InitializeCache(cfg *config.Config) {
	RedisVault = redis.NewClient(&redis.Options{
		Addr:     cfg.REDISHOST,
		Password: "", // no password set
		DB:       0,  // use default DB
		Protocol: 3,  // specify 2 for RESP 2 or 3 for RESP 3
	})
}
