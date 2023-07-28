package generics

import "fmt"


func add[T int | float64](args ...T) T {
	var total T
	for _, v := range(args) {
		total += v
	}
	return total
}


func Main() {
	floatResult := add(1.23, 13.14, 6.485)
	fmt.Printf("Result of the sum of float parameters: %f\n", floatResult)

	integerResult := add(1, 2, 3, 4, 5)
	fmt.Printf("Result of the sum of integer parameters: %d\n", integerResult)

	sliceParams := []int{10, 20, 30}
	sliceParamsResult := add(sliceParams...)
	fmt.Printf("Result of the sum of slice parameters: %d\n", sliceParamsResult)

}