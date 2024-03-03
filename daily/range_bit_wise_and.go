package main

import (
	"fmt"
)

func rangeBitwiseAnd(left int, right int) int {
	count := 0
	for left != right {
		left >>= 1
		right >>= 1
		count++
	}
	return left << count
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
