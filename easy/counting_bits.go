package main

import "fmt"

func dfFunc(n int) int {
	count := 0
	for n > 0 {
		if n%2 == 1 {
			count++
		}

		n = n / 2
	}

	return count
}

func countBits(n int) []int {
	df := make([]int, n+1)

	for i := 1; i <= n; i++ {
		df[i] = dfFunc(i)
	}

	return df
}

func main() {
	fmt.Println(3 & 1)
}
