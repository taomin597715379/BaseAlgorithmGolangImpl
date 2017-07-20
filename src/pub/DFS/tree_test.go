package pub

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	r := createTree()
	fmt.Println(maxDepth(r.root))
	fmt.Println(r.find(4).getData())
}

func Test_allRoadTree(t *testing.T) {
	var s *Stack = NewStack()
	r := createTree()
	allRoadTree(r.root, 9, s)
}

func Test_preOrderR(t *testing.T) {
	r := createTree()
	preOrderR(r.root)
}

func Test_midOrderR(t *testing.T) {
	r := createTree()
	midOrderR(r.root)
}

func Test_postOrderR(t *testing.T) {
	r := createTree()
	postOrderR(r.root)
}

func Test_convertArrayToTree(t *testing.T) {
	var a []int = []int{3, 5, 7, 9, 11, 20}
	r := convertArrayToTree(a, 0, 5)
	midOrderR(r)
}
