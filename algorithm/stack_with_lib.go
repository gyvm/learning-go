package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func outputStack(s *stack.Stack) {
	fmt.Printf("Top ")
	for s.Len() > 0 {
		fmt.Printf("%s->", s.Pop())
	}
	fmt.Printf("End\n")
}

func main() {
	s := stack.New()

	s.Push("A")
	s.Push("B")
	s.Pop()
	s.Push("C")
	s.Push("D")

	outputStack(s)
}

// => Top D->C->A->End
