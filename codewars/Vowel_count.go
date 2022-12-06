package codewars

import (
	"strings"
)

// Return the number (count) of vowels in the given string.

/*We will consider а, e, i, o, u as vowels for this Kata (but not y).*/

// The input string will only consist of lower case letters and/or spaces.

func GetCount(str string) (count int) {
	count = 0
	// a, e, i, o, u
	new_str := strings.Split(str, "")
	for _, v := range new_str {
		if v == "a" || v == "e" || v == "i" || v == "o" || v == "u" {
			count++
		}
	}
	return count
}
