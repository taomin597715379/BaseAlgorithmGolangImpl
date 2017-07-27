package pub

import (
	"errors"
	"fmt"
)

// 能够使用泛型就好了
func max(i, j float64) float64 {
	if i > j {
		return i
	}
	return j
}

func min(i, j float64) float64 {
	if i < j {
		return i
	}
	return j
}

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

// 给一个浮点数序列，取最大乘积连续子串的值，例如 -2.5，4，0，3，0.5，8，-1，则取出的最大乘积连续子串为3，0.5，8。
// 也就是说，上述数组中，3 0.5 8这3个数的乘积30.58=12是最大的，而且是连续的。
func maxProductSubArray1(a []float64) float64 {
	var n int = len(a)
	var maxRet float64 = a[0]
	for i := 0; i < n; i++ {
		var x float64 = 1.0
		for j := i; j < n; j++ {
			x *= a[j]
			if x > maxRet {
				maxRet = x
			}
		}
	}
	return maxRet
}

func maxProductSubArray2(a []float64) float64 {
	var n int = len(a)
	var maxRet float64 = a[0]
	var curMin float64 = a[0]
	var curMax float64 = a[0]
	for i := 1; i < n; i++ {
		x := curMin * a[i]
		y := curMax * a[i]
		curMin = min(min(x, y), a[i])
		curMax = max(max(x, y), a[i])
		maxRet = max(max(maxRet, curMin), curMax)
	}
	return maxRet
}
