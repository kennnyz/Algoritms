package codewars

import (
	"fmt"
	"strings"
)

func Solution(str string) []string {
	if len(str)%2 == 1 {
		str += "_"
	}
	ss := strings.Split(str, "")
	ans := []string{}
	if len(str)%2 == 0 {
		for i := 0; i < len(str)-1; i += 2 {
			ans = append(ans, ss[i]+ss[i+1])

		}
	}
	return ans
}

func main() {

	pol := Solution("123456")
	sol := Solution("1234567")

	fmt.Println(pol)
	fmt.Println(sol)
}
