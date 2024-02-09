package main

import (
	"fmt"
	"math"
)

func minimum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func numSquares(n int) int {
	if n == 1 {
		return 1
	}

	df := make([]int, n+1)
	df[1] = 1

	for i := 2; i <= n; i++ {
		df[i] = df[i-1] + 1

		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			df[i] = minimum(df[i], df[i-j*j]+1)
		}
	}

	return df[n]
}

func main() {
	fmt.Println(numSquares(48))
}
