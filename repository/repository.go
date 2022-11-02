package repository

type Repository interface {
	Set(key string, data interface{}) error
	Get(key string) (string, error)
}

