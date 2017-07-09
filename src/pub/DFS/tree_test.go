package pub

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	r := createTree()
	fmt.Println(maxDepth(r.root))
	fmt.Println(r.Find(4).getData())
}
