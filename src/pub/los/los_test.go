package pub

import (
	"fmt"
	"testing"
)

func Test_balanceArray_1(t *testing.T) {
	var a []int = []int{1, 5, 9, 4}
	var b []int = []int{2, 3, 10, 1}
	ret1, ret2, diff, err := balanceArray(a, b)
	fmt.Println(ret1)
	fmt.Println(ret2)
	fmt.Println(diff)
	fmt.Println(err)
}

func Test_balanceArray_2(t *testing.T) {
	var a []int = []int{2, 7, 1}
	var b []int = []int{1, 5, 7, 6}
	ret1, ret2, diff, err := balanceArray(a, b)
	fmt.Println(ret1)
	fmt.Println(ret2)
	fmt.Println(diff)
	fmt.Println(err)
}

func Test_infinitelyClose(t *testing.T) {
	var b []int = []int{1, 5, 7, 6, 4}
	infinitelyClose(b, 14)
}

func Test_combinationMethod1(t *testing.T) {
	var b []int = []int{10, 5, 1}
	fmt.Println(combinationMethod(b, 100))
}

func Test_combinationMethod2(t *testing.T) {
	var b []int = []int{1, 2}
	fmt.Println(combinationMethod(b, 5))
}

func Test_CalcAllSumPermutation(t *testing.T) {
	var s *Stack = NewStack()
	var b []int = []int{1, 2}
	CalcAllSumPermutation(b, 5, s)
}
