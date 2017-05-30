package pub

import (
	"fmt"
	"testing"
)

func Test_quickSort_1(t *testing.T) {
	var a []int = []int{0, 20, 10, 25, 15, 30, 28, 55, 432, 432432, 4234, 333,
		333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80}
	fmt.Println(a)
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func Test_quickSort_2(t *testing.T) {
	var a []int = []int{0, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
		20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 10}
	fmt.Println(a)
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func Test_quickSort_3(t *testing.T) {
	var a []int = []int{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
		20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20}
	fmt.Println(a)
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func Test_quickSort_4(t *testing.T) {
	var a []int = []int{0}
	fmt.Println(a)
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func Test_quickSort_5(t *testing.T) {
	var a []int = []int{}
	fmt.Println(a)
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
