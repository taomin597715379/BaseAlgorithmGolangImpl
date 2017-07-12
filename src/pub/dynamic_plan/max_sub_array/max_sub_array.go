package pub

import (
	"errors"
	"fmt"
)

// 最大子串和(非连续)
func maxSumArray(a []int) (int, error) {
	if len(a) == 0 {
		return 4294967295, errors.New("lenght of array is zero")
	}
	if len(a) == 1 {
		return a[0], nil
	}
	var maxSum int = a[0]
	var currSum int = 0
	for _, v := range a {
		if currSum > 0 {
			currSum += v
		} else {
			currSum = v
		}
		if maxSum < currSum {
			maxSum = currSum
		}
	}
	return maxSum, nil
}

// 最大连续子串和
func maxSumSerialArray(a []int) (int, error) {
	var start, end int
	if len(a) == 0 {
		return 4294967295, errors.New("lenght of array is zero")
	}
	if len(a) == 1 {
		return a[0], nil
	}
	var maxSum int = a[0]
	var currSum int = 0
	for k, v := range a {
		if currSum+v <= v {
			currSum = v
			start = k
		} else {
			currSum += v
		}
		if maxSum < currSum {
			maxSum = currSum
			end = k
		}
	}
	fmt.Print(a[start : end+1])
	return maxSum, nil
}
