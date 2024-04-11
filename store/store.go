package store

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type RedisConn interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
}

func GetCurrentBlock(ctx context.Context, r RedisConn, chainId uint64, version uint64, startBlock uint64) (uint64, error) {
	rKey := "chain:" + strconv.FormatUint(chainId, 10) + ":" + strconv.FormatUint(version, 10)

	val, err := r.Get(ctx, rKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			if err := r.Set(ctx, rKey, startBlock, 0).Err(); err != nil {
				return 0, err
			}

			return startBlock, nil
		}

		return 0, err
	}

	ret, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func UpdateCurrentBlock(ctx context.Context, r RedisConn, chainId uint64, version uint64, block uint64) error {
	rKey := "chain:" + strconv.FormatUint(chainId, 10) + ":" + strconv.FormatUint(version, 10)
	return r.Set(ctx, rKey, block, 0).Err()
}
