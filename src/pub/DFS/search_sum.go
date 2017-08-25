package pub

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

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

// 一个整数数组，长度为n，将其分为m份，使各份的和相等，求m的最大值。
// 比如{3，2，4，3，6} 可以分成
// {3，2，4，3，6} m=1;
// {3,6}{2,4,3} m=2
// {3,3}{2,4}{6} m=3
// 所以m的最大值为3。
// 为题转换为寻找任意个数，使得和等于sum/i
func equalizationArray(a []int, s *Stack) {
	var sum int
	var n int = len(a)
	for _, v := range a {
		sum += v
	}
	fmt.Println(a)
	for i := 2; i <= n; i++ {
		if sum%i == 0 {
			CalcSumPermutation(a, 0, sum/i, s)
		}
	}
}

// 输入一个正整数数组，将它们连接起来排成一个数，输出能排出的所有数字中最小的一个。例如输入数组{32, 321}，则输出这两个能排成的最小数字32132。
// 这里假设数组里面的数都是不相同的
// 思路: 求数组的所有元素参与的排列组合，然后比较和值，留下最下的排列组合即可,思路比较简单，只要能够正确求出全排列即可
// 这里假设数组排列组合后的最大值也不会超过int类型最大值，如果有越界的情况得将大数变成字符串比较
func numberOfData(k int) float64 {
	return float64(len(strconv.FormatInt(int64(k), 10)))
}

func combinationSum(a []int) float64 {
	var sum float64
	for _, v := range a {
		sum = sum*math.Pow(10, numberOfData(v)) + float64(v)
	}
	return sum
}

func minimalCombinationofArrays(a []int, i int, s *Stack) {
	if i == len(a) {
		f := combinationSum(a)
		if s.Empty() {
			s.Push(f)
		} else {
			if s.Peak().(float64) > f {
				s.Pop()
				s.Push(f)
			}
		}
	} else {
		for j := i; j < len(a); j++ {
			a[j], a[i] = a[i], a[j]
			minimalCombinationofArrays(a, i+1, s)
			a[j], a[i] = a[i], a[j]
		}
	}
}

// 上面的一个变种问题：从n个元素的数组中取出m个数的所有组合,这个问题是从n中数中选择m个数，然后对m个数进行全排列
func CombinationMFromNArrays(a []int, i int, s *Stack) {

}

// N个鸡蛋放到M个篮子中，篮子不能为空，要满足：对任意不大于N的数量，能用若干个篮子中鸡蛋的和表示。
// 写出函数，对输入整数N和M，输出所有可能的鸡蛋的放法。
// 比如对于9个鸡蛋5个篮子 解至少有三组：
// 1 2 4 1 1
// 1 2 2 2 2
// 1 2 3 2 1
// ...
func permutationMfromN(N, n, j, m int, s *Stack) {
	if n == N && j == m {
		s.Print()
	} else if j > m {
		return
	} else if n <= N {
		for i := 1; i <= N; i++ {
			s.Push(i)
			permutationMfromN(N, n+i, j+1, m, s)
			s.Pop()
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

// 思路三：利用快速排序里面的分治思想, 先在数组中随机选择一个数，然后以这个数为分水岭将数组一分为二，
// 那么最小的K个数就存在三种情况：1. 这k个数都在左边的数组里面，2.这k个数都在右边数组里面，3.两个都有
// 时间复杂度O(nlgk)
func random(i, j int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(j-i) + i
}

func QuickSelect(a []int, left, right int) int {
	var i, j, pivot int
	if left < right {
		i, j = left, right
		key := random(left, right)
		pivot = a[key]
		for {
			for a[j] >= pivot && i < j {
				j--
			}
			for a[i] < pivot && i < j {
				i++
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[i], a[key] = a[key], a[i]
	}
	return i
}

func smallestNumberofK3(a []int, k int) []int {
	var n int = len(a)
	var start int = 0
	var end = len(a) - 1
	if n <= k {
		return a
	} else {
		var index int = QuickSelect(a, start, end)
		for k != index {
			if k > index {
				start = index + 1
			} else {
				end = index - 1
			}
			index = QuickSelect(a, start, end)
		}
		return a[:k]
	}
}

// 思路二：先构成一个K大小的堆，然后遍历n-k，同时比较x < kmax， kmax =x 更新堆，时间复杂度O(nlgk)
func smallestNumberofK2(a []int, k int) []int {
	var smallHeap SliceHeap
	if len(a) < k || k <= 0 {
		return []int{}
	}
	smallHeap = a[0:k]
	heap.Init(&smallHeap)
	for i := k; i < len(a); i++ {
		if smallHeap[0] > a[i] {
			heap.Pop(&smallHeap)
			heap.Push(&smallHeap, a[i])
		}
	}
	return smallHeap
}

// 思路一：通过快速排序之后，取前k个数，时间复杂度是O(nlgn)
func partition(a []int, left int, right int) int {
	var pre_index int = left
	var i, j int = left, right
	for {
		for a[i] < a[pre_index] && i < j {
			i++
		}
		for a[j] >= a[pre_index] && j > i {
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[i], a[pre_index] = a[pre_index], a[i]
	return i
}

func quickSort(a []int, left, right int) {
	if left < right {
		m := partition(a, left, right)
		quickSort(a, left, m-1)
		quickSort(a, m+1, right)
	}
}

func smallestNumberofK1(a []int, k int) []int {
	quickSort(a, 0, len(a)-1)
	return a[0:k]
}
