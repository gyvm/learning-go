//　Stack = LIFO構造 (Last In, First Out)

package main

import "fmt"

// Node は Stack のデータ構造の1要素
type Node struct {
	data string
	next *Node
}

// グローバル変数を定義
// Node構造体のメモリアドレスを保持するポインタ。初期値はメモリアドレスが設定されていない状態
var top *Node = nil
var size int

func push(element string) {
	if top == nil {
		top = new(Node)
		top.data = element
	} else {
		var newNode *Node = new(Node)
		newNode.data = element
		newNode.next = top
		top = newNode
	}
	size++
}

func pop() *Node {
	if top == nil {
		return nil
	}

	// popされるNodeはいつメモリ上から削除されるのか？
	// このケースでは、`return p` しているので、外部から参照される可能性があるので保持される
	// `return p` されない場合や、Nodeへの参照がなくなった時点でガーベジコレクションの対象になる
	var p = top
	top = top.next

	p.next = nil
	size--
	return p
}

func output() {
	fmt.Printf("Top ")
	var node *Node = nil
	for {
		// 最後に追加したNodeのnextのNodeを取得している
		// nextのNodeが存在しない場合は、繰り返し終了
		node = pop()
		if node == nil {
			break
		}
		fmt.Printf("%s->", node.data)
	}
	fmt.Printf("End\n")
}

func main() {
	push("A")
	push("B")
	pop()
	push("C")
	push("D")

	output()
} // => Top D->C->A->End
