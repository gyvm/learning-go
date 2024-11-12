package main

import (
	"fmt"
)

type Node struct {
	data string
	prev *Node
	next *Node
}

var head *Node = nil
var tail *Node = nil
var size int

// キューに要素を追加する
func offer(element string) {
	if head == nil {
		head = new(Node)
		head.data = element
		tail = head
	} else {
		var newNode = new(Node)
		newNode.data = element
		newNode.prev = tail
		tail.next = newNode
		tail = newNode
	}
	size++
}

// キューから要素を取り出す
func poll() *Node {
	if head == nil {
		return nil
	}

	var p = head
	head = head.next

	if head == nil {
		tail = nil
	} else {
		head.prev = nil
	}

	// 取り出したNodeの前後関係は不要なので削除
	p.next = nil
	p.prev = nil
	size--

	return p
}

func output() {
	fmt.Printf("Head<-")
	var node *Node = nil
	for {
		node = poll()
		if node == nil {
			break
		}
		fmt.Printf("%s<-", node.data)
	}
	fmt.Printf("Tail\n")
}

func main() {
	offer("A")
	offer("B")
	offer("C")
	offer("D")

	output()
}

// => Head<-A<-B<-C<-D<-Tail
