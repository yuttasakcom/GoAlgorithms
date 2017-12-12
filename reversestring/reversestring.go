package reversestring

import (
	"strings"
)

func reverseString(s string) string {
	reversed := ""

	chars := strings.Split(s, "")

	for _, char := range chars {
		reversed = char + reversed
	}

	return reversed
}
