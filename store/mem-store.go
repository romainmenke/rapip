package store

import (
	"context"
	"sync"
)

type MemStore struct {
	sync.RWMutex

	inBound chan data
	store   map[string]map[string]interface{}
}

func NewMemStore(ctx context.Context) Store {
	s := &MemStore{
		inBound: make(chan data, 100),
		store:   make(map[string]map[string]interface{}),
	}

	go s.loop(ctx)

	return s
}

func (s *MemStore) Get(ctx context.Context, path string) map[string]interface{} {
	s.RLock()
	defer s.RUnlock()

	vv, ok := s.store[path]
	if !ok {
		return make(map[string]interface{})
	}

	vc := make(map[string]interface{}, len(vv))
	for k, v := range vv {
		vc[k] = v
	}

	return vc
}

func (s *MemStore) Put(ctx context.Context, path string, key string, value interface{}) error {
	d := data{
		path:  path,
		key:   key,
		value: value,
	}

	select {
	case s.inBound <- d:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}

func (s *MemStore) loop(ctx context.Context) {
	for {
		select {
		case d := <-s.inBound:

			s.put(d)

		case <-ctx.Done():
			return
		}
	}
}

func (s *MemStore) put(d data) {
	s.Lock()
	defer s.Unlock()

	vv, ok := s.store[d.path]
	if !ok {
		pathMap := make(map[string]interface{})
		s.store[d.path] = pathMap
		vv = pathMap
	}

	vv[d.key] = d.value
}
