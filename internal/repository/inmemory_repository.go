package repository

import (
	"fmt"
	"github.com/recoilme/pudge"
)

type InMemoryRepository struct {
	db *pudge.Db
}

func NewInMemoryRepository(db *pudge.Db) *InMemoryRepository {
	return &InMemoryRepository{
		db: db,
	}
}

func (pr *InMemoryRepository) Set(key string, url string) error {
	has, err := pr.db.Has(key)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("key is already exist")
	}

	err = pr.db.Set(key, url)
	if err != nil {
		return err
	}

	return nil
}

func (pr *InMemoryRepository) Get(key string) (string, error) {
	var result string
	if err := pr.db.Get(key, &result); err != nil {
		return "", err
	}
	return result, nil
}
