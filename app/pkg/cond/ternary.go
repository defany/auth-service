package cond

func Ternary[T any](cond bool, t T, f T) T {
	if cond {
		return t
	}

	return f
}
