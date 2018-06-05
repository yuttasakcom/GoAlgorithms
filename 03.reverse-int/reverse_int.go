package reverseint

import (
	"fmt"
	"math"
	"strconv"
)

func reverseInt(n int) int {

	sign := math.Signbit(float64(n))

	intString := strconv.Itoa(int(math.Abs(float64(n))))

	r := []rune(intString)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	newInt, err := strconv.Atoi(string(r))

	if err != nil {
		fmt.Println("Error converting string to int")
	}

	if sign == true {
		return newInt * -1
	}

	return newInt
}
