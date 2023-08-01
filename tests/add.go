package add

type NumType interface {
	int | float64
}

func Add[T NumType](args ...T) T {
	var total T
	for _, value := range (args) {
		total += value
	}
	return total
}