package codewars

import (
	"strings"
)

func DNAStrand(dna string) string {
	// your code here
	new_shalom := strings.Split(dna, "")

	strop := ""
	for _, v := range new_shalom {
		if v == "A" {
			v = "T"

		} else if v == "T" {
			v = "A"
		}
		if v == "C" {
			v = "G"

		} else if v == "G" {
			v = "C"
		}

		strop += v
	}
	return strop
}
