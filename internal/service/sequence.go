package service

import "sync/atomic"

type dataSource interface {
	GetLastIndex() uint64
}

type sequence struct {
	currentNumber atomic.Uint64
}

func (s *sequence) setStartNumber(d dataSource) {
	index := d.GetLastIndex()
	if index != 0 {
		s.currentNumber.Store(index + 1)
	}
}

func (s *sequence) next() uint64 {
	result := s.currentNumber.Load()
	s.currentNumber.Add(1)
	return result
}
