package cosmicmodel

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type CosmosDB struct {
	DB          *gorm.DB
	RedisClient *redis.Client
	Lock        sync.Mutex
	ctx         context.Context
}

func DivineNavel(db *gorm.DB, redisClient *redis.Client, ipPool string) *CosmosDB {
	ctx := context.Background()

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return nil
	}

	return &CosmosDB{
		DB:          db,
		RedisClient: redisClient,
		ctx:         context.Background(),
	}
}
