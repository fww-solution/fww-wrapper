package redis

import (
	"context"
	"errors"

	"github.com/go-redsync/redsync/v4"
	redispool "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	redis "github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

// InitRedisClient initializes the Redis client.
func InitRedisClient(client *redis.Client) {
	redisClient = client
}

// CacheMiddleware is a middleware that checks if the response is already cached in Redis.
func CacheMiddleware(c *fiber.Ctx) error {
	// Check if the response is already cached in Redis
	cachedResponse, err := redisClient.Get(ctx, "cached_response").Result()
	if err == nil {
		// If the response is cached, return it
		return c.SendString(cachedResponse)
	}

	// Continue to the next middleware or route
	return c.Next()
}

// Distributed locking with redis
// https://redis.io/topics/distlock

// InitMutex initializes a mutex with the given name.
func InitMutex(mutexName string) *redsync.Mutex {
	pool := goredis.NewPool(redisClient)
	rs := redsync.New(pool)
	mutex := rs.NewMutex(mutexName)
	return mutex
}

// LockMutex locks the mutex. It returns an error if the mutex is already locked on the calling process.
func LockMutex(pool redispool.Pool, mutex *redsync.Mutex) error {
	if err := mutex.Lock(); err != nil {
		return errors.New("failed to lock")
	}
	return nil
}

// UnlockMutex unlocks the mutex. It returns an error if the mutex is not locked on the calling process.
func UnlockMutex(pool redispool.Pool, mutex *redsync.Mutex) error {
	if ok, err := mutex.Unlock(); !ok || err != nil {
		return errors.New("failed to unlock")
	}
	return nil
}
