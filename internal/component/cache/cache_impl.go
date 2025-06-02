package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
	"github.com/vmihailenco/msgpack/v4"
)

type CacheImpl struct {
	redisClient *redis.Client
}

func NewCache(redisClient *redis.Client) *CacheImpl {
	return &CacheImpl{redisClient: redisClient}
}

func (c *CacheImpl) Ping(ctx context.Context) error {
	_, err := c.redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *CacheImpl) CreateRawString(ctx context.Context, key string, value string, expiration time.Duration) (bool, error) {
	// Store the string value directly without serialization or compression
	err := c.redisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		log.Errorf("Failed to save raw string on Redis: %s # err %s", key, err)
		return false, err
	}

	return true, nil
}

// Add this method to get raw strings
func (c *CacheImpl) GetRawString(ctx context.Context, key string) (string, error) {
	// Get the string value directly
	val, err := c.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		log.Errorf("Failed to get raw string from Redis: %s # err %s", key, err)
		return "", err
	}

	return val, nil
}

func (c *CacheImpl) Create(ctx context.Context, key string, data interface{}, expiration time.Duration) (bool, error) {

	serializedData, err := msgpack.Marshal(&data)
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}

		log.Errorf("Failed for marshaling data:", err)
		return false, err
	}

	//compress data:
	compressedData, err := CompressData(serializedData)

	err = c.redisClient.Set(ctx, key, compressedData, expiration).Err()
	if err != nil {
		log.Errorf("Failed save data on Redis: %s # err %s", key, err)
		return false, nil
	}

	return true, nil
}

func (c *CacheImpl) Get(ctx context.Context, key string, data interface{}) (interface{}, error) {

	redisData, err := c.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		log.Errorf("Failed get data from Redis:", err)
		return nil, err
	}

	//decompress data:
	decompressedData, err := DecompressData(redisData, len(redisData))

	err = msgpack.Unmarshal(decompressedData, data)
	if err != nil {
		log.Errorf("Failed for unMarshaling data:", err)
		return nil, err
	}

	return data, nil
}

func (c *CacheImpl) Delete(ctx context.Context, key string) (bool, error) {
	_, err := c.redisClient.Del(ctx, key).Result()
	if err != nil {
		log.Errorf("Failed for delete data on redis:", err)
		return false, err
	}

	return true, nil
}

func (c *CacheImpl) GetOrCreate(
	ctx context.Context,
	key string,
	expiration time.Duration,
	factory func() (interface{}, error),
	data interface{},
) (interface{}, error) {

	cached, err := c.Get(ctx, key, data)
	if err != nil {
		return nil, err
	}

	if cached == nil {

		result, err := factory()
		if err != nil {
			return nil, err
		}

		success, err := c.Create(ctx, key, result, expiration)
		if err != nil || !success {
			log.Errorf("Failed save data on Redis:", err)
		}

		return result, nil
	}

	return cached, nil
}

func (r *CacheImpl) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	// Marshal the value to JSON
	jsonData, err := json.Marshal(value)
	if err != nil {
		return false, fmt.Errorf("failed to marshal cache data: %w", err)
	}

	// Use the Redis SetNX command
	result, err := r.redisClient.SetNX(ctx, key, jsonData, expiration).Result()
	if err != nil {
		return false, fmt.Errorf("failed to set value with NX option: %w", err)
	}

	return result, nil
}

// CreateWithNX creates a key-value pair only if the key doesn't exist
// This is similar to SetNX but specifically for creating new entries
func (c *CacheImpl) CreateWithNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	// Serialize the value
	serializedData, err := msgpack.Marshal(&value)
	if err != nil {
		return false, err
	}

	// Compress data
	compressedData, err := CompressData(serializedData)
	if err != nil {
		return false, err
	}

	// Use SetNX to create only if not exists
	result, err := c.redisClient.SetNX(ctx, key, compressedData, expiration).Result()
	if err != nil {
		return false, err
	}

	return result, nil
}

// UpdateWithOptimisticLock updates a key using optimistic locking to prevent race conditions
// It uses Redis WATCH/MULTI/EXEC to ensure atomic updates
func (c *CacheImpl) UpdateWithOptimisticLock(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	// Maximum number of retries for optimistic locking
	maxRetries := 3
	var err error

	// Serialize the value outside the retry loop
	serializedData, err := msgpack.Marshal(&value)
	if err != nil {
		return false, err
	}

	// Compress data
	compressedData, err := CompressData(serializedData)
	if err != nil {
		return false, err
	}

	for attempt := 0; attempt < maxRetries; attempt++ {
		// Start a transaction with WATCH
		txf := func(tx *redis.Tx) error {
			// Check if the key exists first
			_, err := tx.Get(ctx, key).Result()
			if err != nil && err != redis.Nil {
				return err
			}

			// Create the transaction pipeline
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				// Set the new value with expiration
				pipe.Set(ctx, key, compressedData, expiration)
				return nil
			})
			return err
		}

		// Execute the transaction with optimistic locking
		err = c.redisClient.Watch(ctx, txf, key)
		if err == nil {
			// Transaction succeeded
			return true, nil
		}

		if err == redis.TxFailedErr {
			// Optimistic lock failed, retry after a short delay with exponential backoff
			backoff := time.Duration(20*(1<<attempt)) * time.Millisecond
			jitter := time.Duration(rand.Intn(20)) * time.Millisecond
			time.Sleep(backoff + jitter)
			continue
		}

		// Other error occurred
		return false, err
	}

	// All retries failed
	return false, err
}

// RawUpdateWithOptimisticLock updates a key using optimistic locking without data compression
// It uses Redis WATCH/MULTI/EXEC to ensure atomic updates
func (c *CacheImpl) RawUpdateWithOptimisticLock(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	// Maximum number of retries for optimistic locking
	maxRetries := 3
	var err error

	// Serialize the value outside the retry loop
	serializedData, err := msgpack.Marshal(&value)
	if err != nil {
		return false, err
	}

	for attempt := 0; attempt < maxRetries; attempt++ {
		// Start a transaction with WATCH
		txf := func(tx *redis.Tx) error {
			// Check if the key exists first
			_, err := tx.Get(ctx, key).Result()
			if err != nil && err != redis.Nil {
				return err
			}

			// Create the transaction pipeline
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				// Set the new value with expiration
				pipe.Set(ctx, key, serializedData, expiration)
				return nil
			})
			return err
		}

		// Execute the transaction with optimistic locking
		err = c.redisClient.Watch(ctx, txf, key)
		if err == nil {
			// Transaction succeeded
			return true, nil
		}

		if err == redis.TxFailedErr {
			// Optimistic lock failed, retry after a short delay with exponential backoff
			backoff := time.Duration(20*(1<<attempt)) * time.Millisecond
			jitter := time.Duration(rand.Intn(20)) * time.Millisecond
			time.Sleep(backoff + jitter)
			continue
		}

		// Other error occurred
		return false, err
	}

	// All retries failed
	return false, err
}

// Add this method to return the Redis client directly
func (c *CacheImpl) GetRedisClient() *redis.Client {
	return c.redisClient
}
