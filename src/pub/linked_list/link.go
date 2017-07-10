package pub

import (
	"fmt"
)

type LinkedNode struct {
	Data interface{}
	Next *LinkedNode
}

type Link struct {
	Root *LinkedNode
}

// NewNode
func NewLinkedNode(e interface{}) *LinkedNode {
	return &LinkedNode{Data: e}
}

// NewLink
func NewLinked(ln *LinkedNode) *Link {
	return &Link{Root: ln}
}

// print
func linkPrint(root *LinkedNode) {
	for i := root; i != nil; i = i.Next {
		fmt.Print(i.Data)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}

// need list to test above function
func createTwoLinkedList() (*LinkedNode, *LinkedNode) {
	a1 := NewLinkedNode(1)
	r1 := NewLinked(a1)
	a2 := NewLinkedNode(10)
	r2 := NewLinked(a2)
	a1.Next = NewLinkedNode(2)
	a1.Next.Next = NewLinkedNode(3)
	a1.Next.Next.Next = NewLinkedNode(4)
	a1.Next.Next.Next.Next = NewLinkedNode(5)
	a1.Next.Next.Next.Next.Next = NewLinkedNode(6)
	a2.Next = a1.Next.Next.Next
	a2.Next.Next = NewLinkedNode(7)
	a2.Next.Next.Next = NewLinkedNode(8)
	return r1.Root, r2.Root
}

// need list to test above function
func createOneLinkedList() *LinkedNode {
	a1 := NewLinkedNode(1)
	r := NewLinked(a1)
	a1.Next = NewLinkedNode(2)
	a1.Next.Next = NewLinkedNode(3)
	a1.Next.Next.Next = NewLinkedNode(4)
	a1.Next.Next.Next.Next = NewLinkedNode(5)
	a1.Next.Next.Next.Next.Next = NewLinkedNode(6)
	return r.Root
}

// 利用哈希结构，好处是简单易懂，求解第一次相交的点也比较容易。缺点是得消费O(length(h1))的空间
func isIntersects1(h1 *LinkedNode, h2 *LinkedNode) bool {
	var hashMap = make(map[*LinkedNode]bool, 0)
	for i := h1; i != nil; i = i.Next {
		hashMap[i] = true
	}
	for i := h2; i != nil; i = i.Next {
		if _, ok := hashMap[i]; ok {
			return true
		}
	}
	return false
}

// 如果相交的话，两个链表构成环形。
func isIntersects2(h1 *LinkedNode, h2 *LinkedNode) bool {
	var i *LinkedNode = h1
	var slow, fast *LinkedNode = h1, h1
	for i != nil {
		i = i.Next
	}
	i = h2
	for fast == nil && fast.Next == nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// Title: Define a function, enter a head node of a linked list,
// reverse the list and output the head node of the linked list.
func reverseList(head *LinkedNode) *LinkedNode {
	var p, q, r *LinkedNode
	if head == nil || head.Next == nil {
		return head
	}
	p = head
	q = head.Next
	head.Next = nil
	for q != nil {
		r = q.Next
		q.Next = p
		p = q
		q = r
	}
	head = p
	return head
}

// 倒数第n个元素
func lastK(k int, head *LinkedNode) *LinkedNode {
	var s int
	var i, j *LinkedNode = head, head
	if head == nil {
		return nil
	}
	for i != nil {
		if s == k {
			break
		}
		s++
		i = i.Next
	}
	if s < k {
		return nil
	}
	for i != nil {
		i = i.Next
		j = j.Next
	}
	return j
}
