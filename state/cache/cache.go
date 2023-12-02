package cache

type Cache interface {
	WriteHistoryCache(p1, p2 string)
	CheckHistoryExistCache(p1 string) (bool, string)
}
