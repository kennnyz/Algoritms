package main

import (
	"fmt"

	"github.com/kennyz/Algoritms/leetcode"
)

func main() {
	a := []int{-4, -1, 0, 3, 10}
	leetcode.Rotate(a, 1)
	fmt.Println(a)
}
