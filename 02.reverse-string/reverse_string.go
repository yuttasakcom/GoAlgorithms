package reversestring

func reverseString(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

// Concat
// func reverseString(s string) string {
// 	reversed := ""

// 	chars := strings.Split(s, "")

// 	for _, char := range chars {
// 		reversed = char + reversed
// 	}

// 	return reversed
// }
