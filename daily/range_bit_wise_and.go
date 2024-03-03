package main

import (
	"fmt"
	"math"
)

func rangeBitwiseAnd(left int, right int) int {
	if left == 0 {
		return 0
	}
	if left == right {
		return left
	}

	var logL, logR int

	logL = int(math.Log2(float64(left)))
	logR = int(math.Log2(float64(right)))

	if logL == logR {
		and := left
		for i := left + 1; i <= right; i += 2 {
			and &= i
		}

		return and
	}

	return 0
}

func main() {
	left := 8
	right := 24
	count := 0
	for left != right {
		fmt.Println(left, right)
		left >>= 1
		right >>= 1
		count += 1
	}

	fmt.Println(left << count)
}
