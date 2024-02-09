package main

import "fmt"

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := range res {
		res[i] = make([]int, i+1)
	}

	res[0][0] = 1

	for i := 1; i < numRows; i++ {
		n := len(res[i])
		res[i][0] = 1
		res[i][n-1] = 1
		for j := 1; j < n/2+n%2; j++ {

			v := res[i-1][j-1] + res[i-1][j]

			res[i][j] = v
			res[i][n-j-1] = v
		}
	}

	return res
}

func main() {
	fmt.Println(generate(5))
}
