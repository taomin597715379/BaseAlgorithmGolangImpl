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

func Test_equalizationArray(t *testing.T) {
	var s *Stack = NewStack()
	var a []int = []int{3, 2, 4, 3, 6}
	equalizationArray(a, s)
}

func Test_minimalCombinationofArrays(t *testing.T) {
	var s *Stack = NewStack()
	var a []int = []int{3, 32, 321}
	minimalCombinationofArrays(a, 0, s)
	fmt.Println(s.Peak().(float64))
}

func Test_CalcAllSumPermutation(t *testing.T) {
	var s *Stack = NewStack()
	var a []int = []int{1, 3, 2}
	CalcAllSumPermutation(a, 4, s)
}

func Test_smallestNumberofK1(t *testing.T) {
	var a []int = []int{5, 6, 0, 2, 1, 7, 3}
	fmt.Println(smallestNumberofK1(a, 3))
}

func Test_smallestNumberofK2(t *testing.T) {
	var a []int = []int{5, 6, 0, 2, 1, 7, 3}
	fmt.Println(smallestNumberofK2(a, 3))
}

func Test_smallestNumberofK3(t *testing.T) {
	var a []int = []int{5, 6, 0, 2, 1, 7, 3}
	fmt.Println(smallestNumberofK3(a, 4))
}
