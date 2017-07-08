package pub

import (
	"fmt"
	"testing"
)

func Test_CalcTwoPermutation(t *testing.T) {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(CalcTwoPermutation(a, 10))
}

func Test_CalcThreePermutation_1(t *testing.T) {
	var a []int = []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	CalcThreePermutation(a, 0)
}

func Test_CalcThreePermutation_2(t *testing.T) {
	var a []int = []int{1, 2, -1, 3, 4, 5}
	CalcThreePermutation(a, 6)
}

func Test_CalcSumPermutation(t *testing.T) {
	var s *Stack = NewStack()
	var a []int = []int{1, 3, 2, 4, 6}
	CalcSumPermutation(a, 0, 10, s)
}

func Test_CalcAllSumPermutation(t *testing.T) {
	var s *Stack = NewStack()
	var a []int = []int{1, 3, 2}
	CalcAllSumPermutation(a, 4, s)
}
