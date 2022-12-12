package lib

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

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

func AbsDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Int(s string) int {
	d, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return d
}
func Contains[T comparable](s []T, n T) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}
