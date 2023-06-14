package compose

func Compose[T, U, V any](outer func(U) V, inner func(T) U, x T) V {
	return outer(inner(x))
}
