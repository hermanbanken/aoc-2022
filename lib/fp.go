package lib

func Map[T any, R any](ts []T, fn func(T) R) (out []R) {
	out = make([]R, len(ts))
	for i, t := range ts {
		out[i] = fn(t)
	}
	return
}

func Filter[T any](ts []T, fn func(T) bool) (out []T) {
	out = make([]T, 0, len(ts))
	for _, t := range ts {
		if fn(t) {
			out = append(out, t)
		}
	}
	return
}

func Stride[T any](slice []T, step int, size int, fn func(part []T, startIndex int, list []T)) {
	for i := 0; i < len(slice)-size; i += step {
		fn(slice[i:i+size], i, slice)
	}
}

func StrideMap[T any, R any](slice []T, step int, size int, fn func(part []T, startIndex int, list []T) R) (out []R) {
	for i := 0; i < len(slice)-size; i += step {
		out = append(out, fn(slice[i:i+size], i, slice))
	}
	return
}
