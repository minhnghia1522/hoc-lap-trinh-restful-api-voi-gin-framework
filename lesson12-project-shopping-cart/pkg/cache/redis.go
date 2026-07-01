package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCacheService struct {
	ctx context.Context
	rdb *redis.Client
}

func NewRedisCacheService(rdb *redis.Client) RedisCacheService {
	return &redisCacheService{
		ctx: context.Background(),
		rdb: rdb,
	}
}

func (r *redisCacheService) Set(key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.rdb.Set(r.ctx, key, data, ttl).Err()
}

func (r *redisCacheService) Get(key string, dest any) error {
	val, err := r.rdb.Get(r.ctx, key).Result()
	if err == redis.Nil || err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(val), dest); err != nil {
		return err
	}
	return nil
}

func (r *redisCacheService) Delete(key string) error {
	return r.rdb.Del(r.ctx, key).Err()
}

func (cs *redisCacheService) Clear(pattern string) error {
	cursor := uint64(0)
	for {
		keys, nextCursor, err := cs.rdb.Scan(cs.ctx, cursor, pattern, 2).Result()
		if err != nil {
			return err
		}

		if len(keys) > 0 {
			cs.rdb.Del(cs.ctx, keys...)
		}

		cursor = nextCursor

		if cursor == 0 {
			break
		}
	}

	return nil
}

func (cs *redisCacheService) Exists(key string) (bool, error) {
	count, err := cs.rdb.Exists(cs.ctx, key).Result()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
