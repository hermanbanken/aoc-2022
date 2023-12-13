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

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Identity[T any](in T) T {
	return in
}

func Top[T any, R constraints.Ordered](in []T, fn func(T) R) R {
	if len(in) == 0 {
		var zero R
		return zero
	}
	max := fn(in[0])
	for i := 1; i < len(in); i++ {
		item := fn(in[i])
		if item > max {
			max = item
		}
	}
	return max
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

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func Unit(a int) int {
	if a < 0 {
		return -1
	}
	return 1
}

func Int(s string) int {
	d, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return d
}
func Int64(s string) int64 {
	d, err := strconv.ParseInt(s, 10, 64)
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
