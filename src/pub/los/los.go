package pub

import (
	"errors"
	"fmt"
)

func sumArray(a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

func balanceArray(a, b []int) ([]int, []int, int, error) {
	var an int = len(a)
	var bn int = len(b)
	var isContinue bool = true
	if an != bn {
		return []int{}, []int{}, 0, errors.New(`Incorrect Input Array ...`)
	}
	var currentDiff int = sumArray(a) - sumArray(b)
	if currentDiff < 0 {
		t := a
		a = b
		b = t
	}
	for isContinue {
		isContinue = false
		for i := 0; i < an; i++ {
			for j := 0; j < an; j++ {
				var eDiff int = a[i] - b[j]
				currentDiff = sumArray(a) - sumArray(b)
				if eDiff < currentDiff && eDiff > 0 {
					a[i], b[j] = b[j], a[i]
					isContinue = true
				}
			}
		}
	}
	if currentDiff < 0 {
		currentDiff = -1 * currentDiff
	}
	return a, b, currentDiff, nil
}
