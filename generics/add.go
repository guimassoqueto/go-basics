package generics

import "fmt"


func add[T int | float64](args ...T) T {
	var total T
	for _, v := range(args) {
		total += v
	}
	return total
}

// Generics Constrait
type NumbersTypes interface {
	float64 | int
}
func addV2[T NumbersTypes](args ...T) T {
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


	floatResultV2 := addV2(1.23, 13.14, 6.485)
	fmt.Printf("Result of the sum of float parameters (version 2): %f\n", floatResultV2)

	integerResultV2 := addV2(1, 2, 3, 4, 5)
	fmt.Printf("Result of the sum of integer parameters (version 2): %d\n", integerResultV2)

	sliceParams2 := []int{-1, -7, 20}
	sliceParamsResultV2 := addV2(sliceParams2...)
	fmt.Printf("Result of the sum of slice parameters (version 2): %d\n", sliceParamsResultV2)

}