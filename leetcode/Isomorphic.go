// 205. Isomorphic Strings
// Given two strings s and t, determine if they are isomorphic.
//Two strings s and t are isomorphic if the characters in s can be replaced to get t.

package leetcode

import (
	"strings"
)

func IsIsomorphic(s, t string) bool {
	mapST := make(map[string]string)
	mapTS := make(map[string]string)

	ss := strings.Split(s, "")
	tt := strings.Split(t, "")

	for i := 0; i < len(s); i++ {
		if (mapST[ss[i]] != "" && mapST[ss[i]] != tt[i]) || (mapTS[tt[i]] != "" && mapTS[tt[i]] != ss[i]) {
			return false
		}
		mapST[ss[i]] = tt[i]
		mapTS[tt[i]] = ss[i]
	}
	return true
}
