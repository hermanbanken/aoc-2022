package lib

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
