package service

import (
	"errors"
	"fmt"
	"url-shortener/internal/service/repository"
)

const defaultHost = "http://localhost"

//go:generate mockery --name Repository
type Repository interface {
	Set(id uint64, url string) error
	Get(id uint64) (string, error)
}

type Shortener struct {
	repositories []Repository
	seq          sequence
	addr         string
}

func New(port string, repositories ...Repository) *Shortener {
	s := &Shortener{
		repositories: repositories,
		addr:         fmt.Sprintf("%s:%s", defaultHost, port),
	}
loop:
	for _, repo := range repositories {
		switch v := repo.(type) {
		case *repository.PostgresRepository:
			s.seq.setStartNumber(v)
			break loop
		}
	}
	return s
}

func (s *Shortener) Shorten(url string) (string, error) {
	id := s.seq.next()
	key := keyById(id)
	result := fmt.Sprint(s.addr + "/" + key)

	for _, repo := range s.repositories {
		if err := repo.Set(id, url); err != nil {
			return "", err
		}
	}

	return result, nil
}

func (s *Shortener) GetOriginal(key string) (string, error) {
	err := errors.New("")
	result := ""
	id := idByKey(key)

	for _, repo := range s.repositories {
		result, err = repo.Get(id)
		if err == nil {
			break
		}
		if err != nil && err != repository.ErrorNotFound {
			return "", err
		}
	}
	if err != nil {
		return "", err
	}

	return result, nil
}
