package git

import (
	"strconv"
	"strings"
)

func RGB(r, g, b int) string {

	var up bool = false
	var down bool = false

	sukk := []int{r, g, b}
	puk := [3]string{}
	for i, v := range sukk {
		if v > 255 && up == false {
			up = true
			v = 255
		}
		if v < 0 && down == false {
			down = true
			v = 0
		}
		puk[i], _ = ConvertInt(strconv.Itoa(v), 10, 16)
	}
	res := ""
	for _, i := range puk {
		if len(i) < 2 {
			i = "0" + i
		}
		res += i
	}
	return strings.ToUpper(res)
}
func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
