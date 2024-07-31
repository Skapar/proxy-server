package proxy

import (
	"sync"
)

type Storage struct {
	data sync.Map
}

func (s *Storage) Store(key string, value interface{}) {
	s.data.Store(key, value)
}

func (s *Storage) Load(key string) (interface{}, bool) {
	return s.data.Load(key)
}