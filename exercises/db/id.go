package db

import "sync/atomic"

type IdGenerator interface {
	next() int64
}

type Sequence struct {
	counter int64
}

func (s *Sequence) next() int64 {
	counter := atomic.AddInt64(&s.counter, 1)
	return counter
}
