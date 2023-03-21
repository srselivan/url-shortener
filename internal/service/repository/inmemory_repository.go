package repository

import (
	"sync"
)

type InMemoryRepository struct {
	db sync.Map
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

func (pr *InMemoryRepository) Set(id uint64, url string) error {
	_, ok := pr.db.Load(id)
	if ok {
		return ErrorAlreadyExists
	}

	pr.db.Store(id, url)

	return nil
}

func (pr *InMemoryRepository) Get(id uint64) (string, error) {
	result := ""

	if value, ok := pr.db.Load(id); !ok {
		return "", ErrorNotFound
	} else {
		result = value.(string)
	}

	return result, nil
}
