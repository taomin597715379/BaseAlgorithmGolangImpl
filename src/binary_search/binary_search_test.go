package binary

import (
	"fmt"
	"testing"
)

func Test_Binary_Search_1(t *testing.T) {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(`Binary_Search_1`, binary_search_1(a, 7, 6))
	var b []int = []int{1}
	fmt.Println(`Binary_Search_1`, binary_search_1(b, 1, 1))
}

func Test_Binary_Search_2(t *testing.T) {
	var a []int = []int{1, 1, 1, 2, 2, 3, 3}
	fmt.Println(`Binary_Search_2`, binary_search_2(a, 7, 3))
	var b []int = []int{1}
	fmt.Println(`Binary_Search_2`, binary_search_2(b, 1, 1))
}

func Test_Binary_Search_3(t *testing.T) {
	var a []int = []int{1, 1, 1, 2, 2, 3, 3, 3}
	fmt.Println(`Binary_Search_3`, binary_search_3(a, 7, 3))
	var b []int = []int{1}
	fmt.Println(`Binary_Search_3`, binary_search_3(b, 1, 1))
	var c []int = []int{3, 3, 3}
	fmt.Println(`Binary_Search_3`, binary_search_3(c, 3, 3))
}

func Test_Binary_Search_4(t *testing.T) {
	var a []int = []int{1, 1, 1, 2, 2, 3, 3, 3}
	fmt.Println(`Binary_Search_4`, binary_search_4(a, 7, 2))
	var b []int = []int{1}
	fmt.Println(`Binary_Search_4`, binary_search_4(b, 1, 1))
	var c []int = []int{1, 3, 3}
	fmt.Println(`Binary_Search_4`, binary_search_4(c, 3, 1))
}

func Test_Binary_Search_5(t *testing.T) {
	var a []int = []int{1, 1, 1, 2, 2, 3, 3, 3}
	fmt.Println(`Binary_Search_5`, binary_search_5(a, 7, 2))
	var b []int = []int{1}
	fmt.Println(`Binary_Search_5`, binary_search_5(b, 1, 1))
	var c []int = []int{1, 3, 3}
	fmt.Println(`Binary_Search_5`, binary_search_5(c, 3, 1))
}
