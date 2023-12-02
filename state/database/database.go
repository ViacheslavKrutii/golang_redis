package database

type DbInterface interface {
	WriteHistory(p1, p2 string)
	CheckHistoryExist(p1 string) (bool, string)
}
