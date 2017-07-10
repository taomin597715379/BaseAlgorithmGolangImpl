package pub

import (
	"fmt"
	"testing"
)

func Test_isIntersects(t *testing.T) {
	h1, h2 := createTwoLinkedList()
	fmt.Println(isIntersects1(h1, h2))
	h3, h4 := createTwoLinkedList()
	fmt.Println(isIntersects2(h3, h4))
}

func Test_reverseList(t *testing.T) {
	h1 := createOneLinkedList()
	h2 := reverseList(h1)
	linkPrint(h2)
}

func Test_lastK_1(t *testing.T) {
	h1 := createOneLinkedList()
	fmt.Println(lastK(2, h1))
}

func Test_lastK_2(t *testing.T) {
	h1 := createOneLinkedList()
	fmt.Println(lastK(1, h1))
}

func Test_lastK_3(t *testing.T) {
	h1 := createOneLinkedList()
	fmt.Println(lastK(10, h1))
}
