package pub

import (
	"errors"
)

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
