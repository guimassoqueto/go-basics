package add

import (
	"testing"
)

type AddDataInt struct {
	args []int
	expectedResult int
}

type AddDataFloat struct {
	args []float64
	expectedResult float64
}

func TestAddInt(t *testing.T) {
	testData:= []AddDataInt {
		{[]int{1,2,3}, 6},
		{[]int{0,-1, -3}, -4},
		{[]int{-1,-2,-3}, -6},
		{[]int{0,-1, 2}, 1},
	}
	
	for _, value := range testData {
		funcResult := Add(value.args...)
		expectResult := value.expectedResult

		if funcResult != expectResult {
			t.Errorf("Add(%v) = %d; expected %d", value.args, funcResult, expectResult)
		}
	}

}

func TestAddFloat(t *testing.T) {
	testData:= []AddDataFloat {
		{[]float64{1.1, 2.23, 3.446}, 6.776},
		{[]float64{0, -2.23, -3.446}, -5.676},
		{[]float64{-1.458, -2.23, -3.446}, -7.13496},
		{[]float64{-1.45896, -2.23, 3.446}, -0.24296},
	}
	
	for _, value := range testData {
		funcResult := Add(value.args...)
		expectResult := value.expectedResult

		if funcResult != expectResult {
			t.Errorf("Add(%v) = %.2f; expected %f", value.args, funcResult, expectResult)
		}
	}

}