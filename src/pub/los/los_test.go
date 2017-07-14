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
