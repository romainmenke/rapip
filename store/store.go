package store

import "context"

type Store interface {
	Get(ctx context.Context, path string) map[string]interface{}
	Put(ctx context.Context, path string, key string, value interface{}) error
}
