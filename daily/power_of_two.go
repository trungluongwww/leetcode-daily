package main

import (
	"fmt"
	"math"
)

func isPowerOfTwo(n int) bool {
	mod := math.Mod(math.Log2(float64(n)), 1)
	return mod == 0
}

func main() {
	for _, t := range testCase {
		fmt.Println("result:", isPowerOfTwo(t.n), "expected:", t.output)
	}
}

var testCase = []struct {
	n      int
	output bool
}{
	{
		n:      1,
		output: true,
	}, {
		n:      16,
		output: true,
	}, {
		n:      3,
		output: false,
	},
}
