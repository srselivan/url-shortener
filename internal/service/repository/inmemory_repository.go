package repository

import (
	"fmt"
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
		return fmt.Errorf("in-memory repository: id is already exist")
	}

	pr.db.Store(id, url)

	return nil
}

func (pr *InMemoryRepository) Get(id uint64) (string, error) {
	result := ""

	if value, ok := pr.db.Load(id); !ok {
		return "", fmt.Errorf("in-memory repository: id is not exist")
	} else {
		result = value.(string)
	}

	return result, nil
}
