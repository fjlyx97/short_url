package services

type UrlConversion interface {
	NextId() (int64, error)
}
