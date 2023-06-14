package group

type Group[E any] []E

func Len[E any](g Group[E]) int {
	return len(g)
}
