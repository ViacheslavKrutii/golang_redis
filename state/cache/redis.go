package cache

import "github.com/redis/go-redis/v9"

type RedisState struct {
	redis *redis.Client
}

func NewRedisState() *RedisState {

	return &RedisState{
		redis.NewClient(&redis.Options{
			Addr: "redis:6379",
			DB:   0,
		})}
}

func (r *RedisState) WriteHistoryCache(p1, p2 string) {

}
