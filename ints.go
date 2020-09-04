package ctxh

import "context"

func Int(ctx context.Context, key interface{}, defaul ...int) (value int) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(int); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Int8(ctx context.Context, key interface{}, defaul ...int8) (value int8) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(int8); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Int16(ctx context.Context, key interface{}, defaul ...int16) (value int16) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(int16); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Int32(ctx context.Context, key interface{}, defaul ...int32) (value int32) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(int32); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Int64(ctx context.Context, key interface{}, defaul ...int64) (value int64) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(int64); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Ints(ctx context.Context, key interface{}, defaul ...[]int) (value []int) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.([]int); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Int8s(ctx context.Context, key interface{}, defaul ...[]int8) (value []int8) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.([]int8); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Int16s(ctx context.Context, key interface{}, defaul ...[]int16) (value []int16) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.([]int16); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Int32s(ctx context.Context, key interface{}, defaul ...[]int32) (value []int32) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.([]int32); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Int64s(ctx context.Context, key interface{}, defaul ...[]int64) (value []int64) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.([]int64); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}