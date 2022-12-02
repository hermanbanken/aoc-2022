package lib

import "golang.org/x/exp/constraints"

func Last[T any](in []T) (out T) {
	if len(in) > 0 {
		return in[len(in)-1]
	}
	return
}

func Sum[T constraints.Integer](in []T) (out T) {
	var sum T
	for _, c := range in {
		sum += c
	}
	return sum
}
