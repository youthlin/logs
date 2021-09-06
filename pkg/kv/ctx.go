package kv

import (
	"context"
	"sync"
)

type kvCtxKeyType struct{}

var kvCtxKey = kvCtxKeyType{}

func Add(ctx context.Context, kvs ...interface{}) context.Context {
	kvMap := ctx.Value(kvCtxKey)
	if kvMap == nil {
		kvMap = &sync.Map{}
		ctx = context.WithValue(ctx, kvCtxKey, kvMap)
	}
	for i := 0; i+1 < len(kvs); i += 2 {
		k := kvs[i]
		v := kvs[i+1]
		kvMap.(*sync.Map).LoadOrStore(k, v)
	}
	return ctx
}

func Get(ctx context.Context) (kvs []interface{}) {
	kvMap := ctx.Value(kvCtxKey)
	if kvMap == nil {
		return nil
	}
	kvMap.(*sync.Map).Range(func(key, value interface{}) bool {
		kvs = append(kvs, key, value)
		return true
	})
	return
}
