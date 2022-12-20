package lib

import "sort"

type Set struct {
	data map[string]interface{}
}

type Stringer interface {
	String() string
}

func (s *Set) Add(v Stringer) {
	for s.data == nil {
		s.data = make(map[string]interface{})
	}
	s.data[v.String()] = v
}

func (g Set) Size() (count int) {
	for range g.data {
		count += 1
	}
	return
}

func Unique[T comparable](t []T, lessFn func(T, T) bool) (out []T) {
	sort.Slice(t, func(i, j int) bool {
		return lessFn(t[i], t[j])
	})
	for i := 0; i < len(t); i++ {
		if i == 0 || t[i] != out[len(out)-1] {
			out = append(out, t[i])
		}
	}
	return
}

func UniqueUsingKey[T comparable](t []T) (out []T) {
	v := map[T]bool{}
	for _, t := range t {
		v[t] = true
	}
	for k := range v {
		out = append(out, k)
	}
	return
}
