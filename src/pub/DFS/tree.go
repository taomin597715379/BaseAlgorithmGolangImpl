package pub

import (
	"fmt"
)

// node
type TreeNode struct {
	Data   interface{}
	parent *TreeNode
	left   *TreeNode
	right  *TreeNode
}

// root
type Tree struct {
	root *TreeNode
}

func (t *Tree) find(e interface{}) *TreeNode {
	if t.root == nil {
		return nil
	}
	p := t.root
	return t.isEqual(e, p)
}

func (t *Tree) isEqual(e interface{}, n *TreeNode) *TreeNode {
	if e == n.getData() {
		return n
	}
	if n.hasLChild() {
		if ln := t.isEqual(e, n.getLChild()); ln != nil {
			return ln
		}
	}
	if n.hasRChild() {
		if ln := t.isEqual(e, n.getRChild()); ln != nil {
			return ln
		}
	}
	return nil
}

func (tn *TreeNode) offParent() {
	if !tn.hasParent() {
		return
	}
	if tn.isLChild() {
		tn.parent.left = nil
	} else {
		tn.parent.right = nil
	}
	tn.parent = nil
}

func (tn *TreeNode) setLChild(lc *TreeNode) *TreeNode {
	olc := tn.left
	if tn.hasLChild() {
		tn.left.offParent()
	}
	if lc != nil {
		lc.offParent()
		tn.left = lc
		lc.setParent(tn)
	}
	return olc
}

func (tn *TreeNode) setRChild(lc *TreeNode) *TreeNode {
	olc := tn.right
	if tn.hasRChild() {
		tn.right.offParent()
	}
	if lc != nil {
		lc.offParent()
		tn.right = lc
		lc.setParent(tn)
	}
	return olc
}

func (tn *TreeNode) getData() interface{} {
	if tn == nil {
		return nil
	}
	return tn.Data
}

func (tn *TreeNode) hasLChild() bool {
	return tn.left != nil
}

func (tn *TreeNode) hasParent() bool {
	return tn.parent != nil
}

func (tn *TreeNode) setParent(n *TreeNode) {
	tn.parent = n
}

func (tn *TreeNode) hasRChild() bool {
	return tn.right != nil
}

func (tn *TreeNode) getLChild() *TreeNode {
	if !tn.hasLChild() {
		return nil
	}
	return tn.left
}

func (tn *TreeNode) getRChild() *TreeNode {
	if !tn.hasRChild() {
		return nil
	}
	return tn.right
}

func (tn *TreeNode) isLChild() bool {
	return tn.hasParent() && tn == tn.parent.left
}

func (tn *TreeNode) isRChild() bool {
	return tn.hasParent() && tn == tn.parent.right
}

func NewTree(root *TreeNode) *Tree {
	return &Tree{root: root}
}

func NewTreeNode(e interface{}) *TreeNode {
	return &TreeNode{Data: e}
}

func createTree() *Tree {
	tn := NewTreeNode(1)
	r := NewTree(tn)
	tn.setLChild(NewTreeNode(2))
	tn.setRChild(NewTreeNode(3))
	tn.getLChild().setRChild(NewTreeNode(4))
	tn.getLChild().setLChild(NewTreeNode(1))
	tn.getLChild().getRChild().setLChild(NewTreeNode(2))
	tn.getRChild().setLChild(NewTreeNode(5))
	tn.getRChild().setRChild(NewTreeNode(6))
	return r
}

// 前序遍历 - 递归
func preOrderR(root *TreeNode) {
	if root != nil {
		fmt.Println(root.getData())
		preOrderR(root.left)
		preOrderR(root.right)
	}
}

// 前序遍历 - 递归
func midOrderR(root *TreeNode) {
	if root != nil {
		midOrderR(root.left)
		fmt.Println(root.getData())
		midOrderR(root.right)
	}
}

// 前序遍历 - 递归
func postOrderR(root *TreeNode) {
	if root != nil {
		postOrderR(root.left)
		postOrderR(root.right)
		fmt.Println(root.getData())
	}
}

//树的最大深度
func maxDepth(root *TreeNode) int {
	var left, right int
	if root == nil {
		return 0
	}
	left = maxDepth(root.left)
	right = maxDepth(root.right)
	if left < right {
		return right + 1
	} else {
		return left + 1
	}
}

//树的所有节点数
func getNodeNumber(root *TreeNode) int {
	var left, right int
	if root == nil {
		return 0
	}
	left = getNodeNumber(root.left)
	right = getNodeNumber(root.right)
	return left + right + 1
}

// 在二元树中找出和为某一值的所有路径
// 输入一个整数和一棵二元树，从树的根结点开始往下访问一直到叶结点所经过的所有结点形成一条路径，然后打印出和与输入整数相等的所有路径。
// 例如输入整数22和如下二元树
func allRoadTree(root *TreeNode, currentSum int, s *Stack) {
	if root == nil {
		return
	}
	s.Push(root.Data.(int))
	var isleaf bool = (root.left == nil) && (root.right == nil)
	if isleaf {
		if currentSum == root.Data.(int) {
			s.Print()
		}
	} else {
		if root.left != nil {
			allRoadTree(root.left, currentSum-root.Data.(int), s)
		}
		if root.right != nil {
			allRoadTree(root.right, currentSum-root.Data.(int), s)
		}
	}
	s.Pop()
}

// 如何将数组转化为二叉搜索树
func convertArrayToTree(a []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	var m int = start + (end-start)/2
	root := NewTreeNode(a[m])
	root.left = convertArrayToTree(a, start, m-1)
	root.right = convertArrayToTree(a, m+1, end)
	return root
}

// 二叉搜索树最近公共祖先LCA(Lowest Common Ancestor)
func getBstLCA(root *TreeNode, x, y int) *TreeNode {
	var t *TreeNode = root
	for {
		if t.getData().(int) < x && t.getData().(int) < y {
			t = t.left
		} else if t.getData().(int) > x && t.getData().(int) > y {
			t = t.right
		} else {
			return t
		}
	}
}
