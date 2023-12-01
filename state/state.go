package state

import (
	"Proj/golang_pub-sub_observe/state/cache"
	"Proj/golang_pub-sub_observe/state/database"
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
