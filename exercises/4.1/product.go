package product

// Your constraint definition goes here!

type Numeric interface {
	~int | ~float64 | ~complex128
}

func Product[T Numeric](x, y T) T {
	return x * y
}
