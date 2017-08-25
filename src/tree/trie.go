package main

import (
	"fmt"
)

type Node struct {
	C     string
	Child []*Node
	Word  string
}

type TrieRoot struct {
	root *Node
}

func NewTrieRoot() *TrieRoot {
	return &TrieRoot{root: &Node{C: ``, Word: ``}}
}

func (tr *TrieRoot) find(node *Node, v string) int {
	child := node.Child
	if len(child) == 0 {
		return -1
	} else {
		for k, cl := range child {
			if cl.C == v {
				return k
			}
		}
		return -1
	}
}

func (tr *TrieRoot) add(word string) {
	node := tr.root
	for _, v := range word {
		pos := tr.find(node, string(v))
		if pos < 0 {
			n := &Node{C: string(v), Word: ``}
			node.Child = append(node.Child, n)
			pos = len(node.Child) - 1
		}
		node = node.Child[pos]
	}
	node.Word = word
}

func (tr *TrieRoot) iterator(node *Node) {
	if node.Word != `` {
		fmt.Println(node.Word)
	}
	for _, v := range node.Child {
		tr.iterator(v)
	}
}

func (tr *TrieRoot) search(node *Node, obj string) {
	var flag bool
	t := node
	if obj == `` {
		fmt.Printf(`Input Is Null...`)
		return
	}
	for _, v := range obj {
		for kl, vl := range t.Child {
			if vl.C == string(v) {
				t = t.Child[kl]
				flag = true
				break
			}
		}
		if flag == false {
			fmt.Printf(`You Search Not Found ... `)
			return
		}
	}
	tr.iterator(t)
}

func AddWord(t *TrieRoot, ws []string) {
	for _, word := range ws {
		t.add(word)
	}
}

func main() {
	var ws = []string{"python", "function", "php", "food", "kiss", "perl", "easy", "goal", "go", "golang"}
	var t *TrieRoot = NewTrieRoot()
	AddWord(t, ws)
	t.search(t.root, `go`)
}
