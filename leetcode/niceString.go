package leetcode

import "fmt"

func isBoring(arr []int) bool {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}
	freqFreq := make(map[int]int)
	for _, val := range freq {
		freqFreq[val]++
	}
	switch len(freqFreq) {
	case 1:
		return true
	case 2:
		var f1, f2 int
		for k, v := range freqFreq {
			if f1 == 0 {
				f1 = k
			} else {
				f2 = k
			}
			if v == 1 {
				return true
			}
		}
		if (f1 == 1 && freqFreq[f1] == 1) || (f2 == 1 && freqFreq[f2] == 1) {
			return true
		}
		if (f1-f2 == 1 && freqFreq[f1] == 1) || (f2-f1 == 1 && freqFreq[f2] == 1) {
			return true
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	l := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isBoring(a[i : j+1]) {
				l = j - i + 1
			}
		}
	}
	fmt.Println(l)
}
