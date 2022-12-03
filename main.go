package main

import (
	"fmt"
)

var pop int = 4

func isBadversion(n int) bool {
	if n >= pop {
		return true
	}
	return false
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(isBadversion(a[3]))
}
