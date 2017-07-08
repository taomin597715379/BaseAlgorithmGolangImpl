package pub

import (
	"container/list"
	"fmt"
)

// 通过循环双向链表构建栈结构，为后续的算法提供一个基础数据结构
type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
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

// 寻找两个数据的和等于给定值,一次即可
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
func CalcTwoPermutation(nums []int, target int) []int {
	var hashmap = make(map[int]int, 0)
	for k, v := range nums {
		hashmap[v] = k
	}
	for k, v := range nums {
		if _, ok := hashmap[target-v]; ok && hashmap[target-v] != k {
			return []int{k, hashmap[target-v]}
		}
	}
	return []int{}
}

// 寻找三个数据的和等于给定值
// 是上面的问题的变种，可以先选中一个数，求定值和该值的差值，然后从剩下的数里面选择出两个数的和等于差值
func CalcThreePermutation(nums []int, target int) {
	var n int = len(nums)
	var hashmap = make(map[int]int, 0)
	for k, v := range nums {
		hashmap[v] = k
	}
	for i := 0; i < n; i++ {
		diff := target - nums[i]
		for k, v := range nums {
			if k != i {
				delete(hashmap, v)
				if _, ok := hashmap[diff-v]; ok && hashmap[diff-v] != k {
					fmt.Println(nums[i], v, nums[hashmap[diff-v]])
				}
				hashmap[v] = k
			}

		}
	}
}

// 寻找任意一个数据的和等于给定值，列出所有组合
// Enter a string that prints all the permutations of the characters in the string.
// For example, enter the string abc, the output by the characters a, b, c can be arranged out of all the strings
// Abc, acb, bac, bca, cab and cba.
func CalcSumPermutation(a []int, i, sum int, s *Stack) {
	if sum == 0 {
		s.Print()
		return
	}
	if i >= len(a) {
		return
	} else {
		s.Push(a[i])
		if sum-a[i] >= 0 {
			CalcSumPermutation(a, i+1, sum-a[i], s)
		}
		s.Pop()
		if sum > 0 {
			CalcSumPermutation(a, i+1, sum, s)
		}
	}
}

// 寻找任意一个数据的和等于给定值，列出所有组合,和上面的问题不同之处在于：一个数字可以无数次出现，只要相加的和等于给定的值
// Enter a string that prints all the permutations of the characters in the string.
// For example, enter the string abc, the output by the characters a, b, c can be arranged out of all the strings
// aaa, aab, aba, baa, abc, ...
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
