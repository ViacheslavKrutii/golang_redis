package state

import (
	"Proj/golang_pub-sub_observe/state/cache"
	"Proj/golang_pub-sub_observe/state/database"
	"log"
)

type StateStruct struct {
	Cache cache.Cache
	Db    database.DbInterface
}

func NewState() *StateStruct {
	return &StateStruct{
		cache.NewRedisState(),
		database.NewMysql(),
	}
}

func (s *StateStruct) WriteHistory(p1, p2 string) {
	s.Cache.WriteHistoryCache(p1, p2)
	s.Db.WriteHistory(p1, p2)
}

func (s *StateStruct) ShowInviteHistory(p1 string) {
	hitoryExist, history := s.Cache.CheckHistoryExistCache(p1)
	if hitoryExist {
		log.Println(history)
	}
	hitoryExist, history = s.Db.CheckHistoryExist(p1)
	if hitoryExist {
		log.Println(history)
	}
}
