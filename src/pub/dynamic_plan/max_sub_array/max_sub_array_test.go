package pub

import (
	"fmt"
	"testing"
)

func Test_maxSumArray_1(t *testing.T) {
	var a []int = []int{0, 20, 10, 25, 15, 30, 28, 55, 432, 432432, 4234, 333,
		333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80}
	b, err := maxSumArray(a)
	if err == nil {
		fmt.Println(b)
	}
}

func Test_maxSumArray_2(t *testing.T) {
	var a []int = []int{-20, -10, 10, -5, 9, 1, -1}
	b, err := maxSumArray(a)
	if err == nil {
		fmt.Println(b)
	}
}

func Test_maxSumArray_3(t *testing.T) {
	var a []int = []int{-20, -10, -10, -5, -9, -1, -1}
	b, err := maxSumArray(a)
	if err == nil {
		fmt.Println(b)
	}
}

func Test_maxSumSerialArray_1(t *testing.T) {
	var a []int = []int{1, -2, 3, 10, -4, 7, 2, -5}
	b, err := maxSumSerialArray(a)
	if err == nil {
		fmt.Println(b)
	}
}

func Test_maxSumSerialArray_2(t *testing.T) {
	var a []int = []int{1, -10, 1, 1, 1}
	b, err := maxSumSerialArray(a)
	if err == nil {
		fmt.Println(b)
	}
}

func Test_insertSort_4(t *testing.T) {
	var a []int = []int{0}
	b, err := maxSumArray(a)
	if err == nil {
		fmt.Println(b)
	}
}

func Test_insertSort_5(t *testing.T) {
	var a []int = []int{}
	b, err := maxSumArray(a)
	if err == nil {
		fmt.Println(b)
	}
}
