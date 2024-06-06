package refs

func P[T any](t T) *T {
	return &t
}
