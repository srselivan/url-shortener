package service

import "sync/atomic"

type sequence struct {
	currentNumber atomic.Uint64
}

func (s *sequence) next() uint64 {
	result := s.currentNumber.Load()
	s.currentNumber.Add(1)
	return result
}
