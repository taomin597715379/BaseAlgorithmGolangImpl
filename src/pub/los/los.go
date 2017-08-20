package pub

import (
	"container/list"
	"errors"
	"fmt"
)

const MAX int = 200

// 通过循环双向链表构建栈结构，为后续的算法提供一个基础数据结构
type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

type SliceHeap []int

func (s SliceHeap) Len() int {
	return len(s)
}

func (s SliceHeap) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s SliceHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
}

func (s *SliceHeap) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *SliceHeap) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func (s *Stack) Pop() interface{} {
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
		return e.Value
	}
	return nil
}

func (s *Stack) Peak() interface{} {
	e := s.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}

func (s *Stack) Print() {
	for e := s.list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Print("\n")
}

func (s *Stack) reWrite(a []interface{}) {
	for e := s.list.Front(); e != nil; e = e.Next() {
		a = append(a, e.Value)
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

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

// 从N个数中选取若干个数之和最接近M
func infinitelyClose(a []int, m int) {
	var dp [MAX]int
	var n int = len(a)
	var i, j int
	for i = 0; i < n; i++ {
		for j = m; j >= a[i]; j-- {
			if j >= a[i] {
				dp[j] = max(dp[j-a[i]]+a[i], dp[j])
			}
		}
	}
	fmt.Println(dp[m : m+1])
}

// 将100这个数，分成10，5，1的组合，会有多少种组合
// 递归解有栈溢出风险
func combinationMethod(a []int, m int) []int {
	var b [MAX]int
	b[0] = 1
	var n int = len(a)
	for i := 0; i < n; i++ {
		for j := a[i]; j <= m; j++ {
			b[j] += b[j-a[i]]
		}
	}
	return b[m : m+1]
}

// 寻找任意一个数据的和等于给定值，列出所有组合,和上面的问题不同之处在于：一个数字可以无数次出现，只要相加的和等于给定的值
// Enter a string that prints all the permutations of the characters in the string.
// For example, enter the string abc, the output by the characters a, b, c can be arranged out of all the strings
// aaa, aab, aba, baa, abc, ...
// 使用递归有个问题是如果数组里的值很小,但要求的和值非常大的时候,很容易出现栈溢出情况
// 还有就是这种方式给出来的结果是有重复方案的，需要去掉
func CalcAllSumPermutation(a []int, sum int, s *Stack) {
	if sum == 0 {
		s.Print()
		return
	} else {
		for i := 0; i < len(a); i++ {
			if sum-a[i] >= 0 {
				s.Push(a[i])
				CalcAllSumPermutation(a, sum-a[i], s)
				s.Pop()
			}
		}
	}

}
