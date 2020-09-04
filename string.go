package ctxh

import "context"

func String(ctx context.Context, key interface{}, defaul ...string) (value string) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(string); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Strings(ctx context.Context, key interface{}, defaul ...[]string) (value []string) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.([]string); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}
