package group

type IdFunc[T any] func(T) T

// but you can't do this:
// type DoesntWorkFunc = func[T any](T) T

func Identity[T any](x T) T {
	return x
}
