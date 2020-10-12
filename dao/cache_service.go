package dao

type CacheInterface interface {
	InitCache(ip string, port string, password string) error
	SetUrl(uidB62 string, url string) error
	GetUrl(url string) (string, error)
}
