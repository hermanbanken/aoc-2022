package lib

// Ternary short-hand
func Ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

// ErrOnly discards any response value except the error
func ErrOnly(_ interface{}, err error) error {
	return err
}

// ErrOnly2 discards any response value except the error
func ErrOnly2(_, _ interface{}, err error) error {
	return err
}

// ErrOnly3 discards any response value except the error
func ErrOnly3(_, _, _ interface{}, err error) error {
	return err
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
