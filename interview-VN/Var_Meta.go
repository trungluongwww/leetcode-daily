package main

import "fmt"

// Problem
// Give an array [0,1,2,3...n] with length is n + 1
// Let's find a missing value if taking an any value in array (length is n)
// Example:
// input: [0,1,2,3,5] -> output: 4

// Solution ...
func Solution(nums []int) int {
	sum := 0
	for i, num := range nums {
		sum += i - num
	}

	return sum + len(nums)
}

func main() {
	fmt.Println(Solution([]int{0, 1, 2, 4, 3, 5}))
}
