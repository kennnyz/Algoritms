package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	freq := make(map[int]int)
	maxPrefixLen := 0
	for i := 0; i < n-1; i++ {
		freq[a[i]]++
		if isBoring(freq, i+1) {
			maxPrefixLen = i + 1
		}
	}

	fmt.Println(maxPrefixLen)
}

func isBoring(freq map[int]int, prefixLen int) bool {
	values := make(map[int]bool)
	for _, freq := range freq {
		values[freq] = true
	}
	if len(values) > 2 {
		return false
	} else if len(values) == 1 {
		return true
	} else {
		var v1, v2 int
		for k := range freq {
			if freq[k] > v1 {
				v2 = v1
				v1 = freq[k]
			} else if freq[k] > v2 {
				v2 = freq[k]
			}
		}
		return (v1-v2 == 1 && v1+v2 == prefixLen) || (v1 == 1 && v2 == prefixLen-1)
	}
}
