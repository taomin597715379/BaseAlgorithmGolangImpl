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
