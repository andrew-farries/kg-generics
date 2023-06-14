package funcmap

type FuncMap[T, U any] map[string]func(T) U

func (m FuncMap[T, U]) Apply(f string, x T) U {
	fn, ok := m[f]
	if !ok {
		panic("function not found")
	}

	return fn(x)
}
