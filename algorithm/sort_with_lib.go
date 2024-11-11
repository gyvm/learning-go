package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{10, 3, 5, 7, 2, 8}
	sort.Ints(numbers)

	fmt.Println(numbers)
}

// => [2 3 5 7 8 10]
