package main

import (
	"fmt"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := map[int]int{}

	for _, num := range arr {
		m[num] += 1
	}

	count := 0

	r := map[int]int{}
	for _, value := range m {
		count += 1
		r[value] += 1
	}

	min := 1

	for k >= min {
		if r[min] != 0 {
			if k/min > r[min] {
				k -= r[min] * min
				count -= r[min]
			} else {
				count -= k / min
				break
			}
		}
		min++
	}

	return count
}

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{2, 2, 1, 5, 3, 3}, 3))
}
