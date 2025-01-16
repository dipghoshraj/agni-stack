package karma

import (
	cosmicmodel "brahma-node-creation-hub/cosmic-model"
	"context"
	"fmt"
	"time"

	"brahma-node-creation-hub/karma/kheersagar"

	"github.com/go-redis/redis/v8"
)

func allocatePort(rdb *redis.Client) (string, error) {
	deadline := time.Now().Add(20 * time.Second) // 20 seconds timeout
	ctx := context.Background()

	for {
		port, err := rdb.SPop(ctx, "available_ports").Result()
		if err == nil {
			return port, nil
		}

		if time.Now().After(deadline) {
			return "", fmt.Errorf("timeout reached, no available ports")
		}

		if err != redis.Nil {
			return "", fmt.Errorf("error fetching port: %w", err)
		}

		fmt.Println("No available ports, retrying...")
		time.Sleep(100 * time.Millisecond) // Small delay before retrying
	}
}

func (ch *CosmosHandler) BengingLok(loka *cosmicmodel.LokaNode) error {

	// tx := ndb.DbManager.DB.Begin()
	tx := ch.CosmosDB.DB.Begin()
	if err := tx.Create(loka).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to store lok details in database: %v", err)
	}

	port, err := allocatePort(ch.CosmosDB.RedisClient)
	if err != nil {
		return fmt.Errorf("failed to port allocation: %v", err)
	}

	err = kheersagar.SpinUpLoka(loka.ID, loka.VolumeName, loka.Capacity, port)
	if err != nil {
		return fmt.Errorf("failed to spin up lok: %v", err)
	}

	loka.Port = port
	if err := tx.Save(loka).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update loka port: %v", err)
	}
	return tx.Commit().Error

}

func (ch *CosmosHandler) GetAllLokas() ([]cosmicmodel.LokaNode, error) {
	var lokas []cosmicmodel.LokaNode
	if err := ch.CosmosDB.DB.Find(&lokas).Error; err != nil {
		return nil, err
	}
	return lokas, nil
}
