package pub

import (
	"fmt"
	"testing"
)

func Test_longIncrSubseq_1(t *testing.T) {
	var a []int = []int{0, 20, 10, 25, 15, 30, 28, 55, 432, 432432, 4234, 333,
		333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80}
	fmt.Println(longIncrSubseq(a))
}

func Test_longIncrSubseq_2(t *testing.T) {
	var a []int = []int{-20, -10, 10, -5, 9, 1, -1}
	fmt.Println(longIncrSubseq(a))
}

func Test_longIncrSubseq_3(t *testing.T) {
	var a []int = []int{-20, -20, -20, -20, -20, -20, -20}
	fmt.Println(longIncrSubseq(a))
}

func Test_longIncrSubseq_4(t *testing.T) {
	var a []int = []int{0}
	fmt.Println(longIncrSubseq(a))
}

func Test_longIncrSubseq_5(t *testing.T) {
	var a []int = []int{}
	fmt.Println(longIncrSubseq(a))
}
