package ctxh

import "context"

func Bool(ctx context.Context, key interface{}, defaul ...bool) (value bool) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(bool); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func BoolOrTrue(ctx context.Context, key interface{}) bool {
	return Bool(ctx, key, true)
}

func BoolOrFalse(ctx context.Context, key interface{}) bool {
	return Bool(ctx, key)
}
