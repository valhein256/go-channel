package main

import (
	"math/rand"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

type Tree struct {
	node *Node
}

func (t *Tree) insert(value int) {
	if t.node == nil {
		println("insert first item to tree", value)
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}
}

func (n *Node) search(value int) int {
	if n == nil {
		return -1
	}

	if value == n.value {
		return n.value
	}

	if value < n.value {
		return n.left.search(value)
	} else {
		return n.right.search(value)
	}
}

func printTree(n *Node, sort_type string) {
	if n == nil {
		return
	}
	switch sort_type {
	case "ascending":
		printTree(n.left, sort_type)
		println(n.value)
		printTree(n.right, sort_type)
		break
	case "descending":
		printTree(n.right, sort_type)
		println(n.value)
		printTree(n.left, sort_type)
		break
	default:
		println(n.value)
		printTree(n.left, sort_type)
		printTree(n.right, sort_type)
		break
	}
}

func process_generate_tree(ch chan *Tree) {
	t := &Tree{}
	for i := 0; i < 10; i++ {
		n := rand.Intn(100)
		t.insert(n)
	}
	ch <- t
}

func main() {
	ch := make(chan *Tree)
	go process_generate_tree(ch)
	t := <-ch

	println("default")
	printTree(t.node, "ascending")
	println()
}
