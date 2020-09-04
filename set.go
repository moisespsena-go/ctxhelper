package ctxh

import (
	"context"
)

type contextValues struct {
	context.Context
	values []map[interface{}]interface{}
}

func (this *contextValues) Value(key interface{}) interface{} {
	for _, m := range this.values {
		if v, ok := m[key]; ok {
			return v
		}
	}
	return this.Context.Value(key)
}

func SetM(ctx context.Context, m ...map[interface{}]interface{}) context.Context {
	if len(m) > 0 {
		return &contextValues{ctx, m}
	}
	return ctx
}
