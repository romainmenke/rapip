package store

import "context"

type data struct {
	path  string
	key   string
	value interface{}
}

type pathKeyType string

const pathKey = pathKeyType("rapip.store.path")

func ContextWithPath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, pathKey, path)
}

func Path(ctx context.Context) (string, bool) {
	v := ctx.Value(pathKey)
	if v == nil {
		return "", false
	}

	s, ok := v.(string)
	if !ok {
		return "", false
	}

	return s, true
}
