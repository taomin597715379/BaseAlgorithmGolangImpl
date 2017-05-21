package pub

import (
	"fmt"
	"testing"
)

func Test_bubbleSort_1(t *testing.T) {
	var a []int = []int{0, 20, 10, 25, 15, 30, 28, 55, 432, 432432, 4234, 333,
		333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80}
	fmt.Println(a)
	b := bubbleSort(a)
	fmt.Println(b)
}

func Test_bubbleSort_2(t *testing.T) {
	var a []int = []int{0, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
		20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 10}
	fmt.Println(a)
	b := bubbleSort(a)
	fmt.Println(b)
}

func Test_bubbleSort_3(t *testing.T) {
	var a []int = []int{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
		20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20}
	fmt.Println(a)
	b := bubbleSort(a)
	fmt.Println(b)
}

func Test_bubbleSort_4(t *testing.T) {
	var a []int = []int{0}
	fmt.Println(a)
	b := bubbleSort(a)
	fmt.Println(b)
}

func Test_bubbleSort_5(t *testing.T) {
	var a []int = []int{}
	fmt.Println(a)
	b := bubbleSort(a)
	fmt.Println(b)
}

func Test_bubbleSortImprv_1(t *testing.T) {
	var a []int = []int{0, 20, 10, 25, 15, 30, 28, 55, 432, 432432, 4234, 333,
		333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80}
	fmt.Println(a)
	b := bubbleSortImprv(a)
	fmt.Println(b)
}

func Test_bubbleSortImprv_2(t *testing.T) {
	var a []int = []int{0, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
		20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 10}
	fmt.Println(a)
	b := bubbleSortImprv(a)
	fmt.Println(b)
}

func Test_bubbleSortImprv_3(t *testing.T) {
	var a []int = []int{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
		20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20}
	fmt.Println(a)
	b := bubbleSortImprv(a)
	fmt.Println(b)
}

func Test_bubbleSortImprv_4(t *testing.T) {
	var a []int = []int{0}
	fmt.Println(a)
	b := bubbleSortImprv(a)
	fmt.Println(b)
}

func Test_bubbleSortImprv_5(t *testing.T) {
	var a []int = []int{}
	fmt.Println(a)
	b := bubbleSortImprv(a)
	fmt.Println(b)
}
