package codewars

import "strings"

func CamelCase(s string) string {
	if len(s) != 0 {
		if s[0] == 32 {
			s = s[1:]
		}
		if s[len(s)-1] == 32 {
			s = s[0 : len(s)-1]
		}

		ss := strings.Split(s, " ")
		ans := ""
		for v := 0; v < len(ss); v++ {
			vv := strings.Split(ss[v], "")
			vv[0] = strings.ToUpper(vv[0])
			new_v := ""
			for _, i := range vv {
				new_v += i
			}
			ans += new_v
		}
		return ans
	}

	return ""

}
