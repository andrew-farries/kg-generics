package intish

// Your Intish interface goes here!

type Intish interface {
	~int
}

func IsPositive[T Intish](v T) bool {
	return v > 0
}
