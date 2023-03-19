package service

import (
	"fmt"
)

type Repository interface {
	Set(id uint64, url string) error
	Get(id uint64) (string, error)
}

type Shortener struct {
	repositories []Repository
	seq          sequence
}

func New(repositories ...Repository) *Shortener {
	return &Shortener{
		repositories: repositories,
	}
}

func (s *Shortener) Shorten(url string) (string, error) {
	id := s.seq.next()
	key := keyById(id)
	result := fmt.Sprint("http://localhost:8080/" + key)

	for _, repo := range s.repositories {
		if err := repo.Set(id, url); err != nil {
			return "", err
		}
	}

	return result, nil
}

func (s *Shortener) GetOriginal(key string) (string, error) {
	result := ""
	id := idByKey(key)

	result, err := s.repositories[0].Get(id)
	if err != nil {
		return "", err
	}

	return result, nil
}
