package ctxh

import "context"

func Uint(ctx context.Context, key interface{}, defaul ...uint) (value uint) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(uint); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Uint8(ctx context.Context, key interface{}, defaul ...uint8) (value uint8) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(uint8); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Uint16(ctx context.Context, key interface{}, defaul ...uint16) (value uint16) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(uint16); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Uint32(ctx context.Context, key interface{}, defaul ...uint32) (value uint32) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(uint32); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Uint64(ctx context.Context, key interface{}, defaul ...uint64) (value uint64) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.(uint64); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Uints(ctx context.Context, key interface{}, defaul ...[]uint) (value []uint) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.([]uint); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Uint8s(ctx context.Context, key interface{}, defaul ...[]uint8) (value []uint8) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.([]uint8); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Uint16s(ctx context.Context, key interface{}, defaul ...[]uint16) (value []uint16) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.([]uint16); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Uint32s(ctx context.Context, key interface{}, defaul ...[]uint32) (value []uint32) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.([]uint32); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}

func Uint64s(ctx context.Context, key interface{}, defaul ...[]uint64) (value []uint64) {
	if v := ctx.Value(key); v != nil {
		var ok bool
		if value, ok = v.([]uint64); ok {
			return
		}
	}

	for _, value = range defaul {
		break
	}
	return
}
