package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

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
	msg := fmt.Sprintf("%v invite %v", p1, p2)
	historyExist, existingHistory := r.CheckHistoryExistCache(p1)
	switch historyExist {
	case true:
		msg = msg + ";" + existingHistory

		err := r.redis.Set(context.TODO(), p1, msg, 0).Err()

		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%s udate history", p1)
	case false:
		err := r.redis.Set(context.TODO(), p1, msg, 0).Err()

		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%s write history", p1)
	}

}

func (r *RedisState) CheckHistoryExistCache(p1 string) (bool, string) {
	existingHistory, err := r.redis.Get(context.TODO(), p1).Result()
	if err != nil {
		panic(err)
	}
	if existingHistory == "" {
		return false, existingHistory
	}
	return true, existingHistory
}
