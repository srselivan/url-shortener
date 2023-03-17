package service

import (
	"fmt"
	"github.com/google/uuid"
)

type Repository interface {
	Set(key string, url string) error
	Get(key string) (string, error)
}

type Shortener struct {
	repositories []Repository
}

func New(repositories ...Repository) *Shortener {
	return &Shortener{
		repositories: repositories,
	}
}

func (s *Shortener) Shorten(url string) (string, error) {
	id := uuid.New().ID()
	key := shorten(uint(id))
	result := fmt.Sprint("http://localhost:8080/" + key)

	for _, repo := range s.repositories {
		if err := repo.Set(key, url); err != nil {
			return "", err
		}
	}

	return result, nil
}

func (s *Shortener) GetOriginal(key string) (string, error) {
	var result string

	result, err := s.repositories[0].Get(key)
	if err != nil {
		return "", err
	}

	return result, nil
}
