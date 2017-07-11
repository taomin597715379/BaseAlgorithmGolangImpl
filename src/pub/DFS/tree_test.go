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
