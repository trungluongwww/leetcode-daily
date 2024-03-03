package main

import "fmt"

func QuickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	i := start
	pivot := end

	for i < pivot {
		if nums[i] > nums[pivot] {
			temp := nums[i]
			nums[i] = nums[pivot-1]
			nums[pivot-1] = nums[pivot]
			nums[pivot] = temp
			pivot--
		} else {
			i++
		}

		fmt.Println(nums)
	}

	QuickSort(nums, start, pivot-1)
	QuickSort(nums, pivot+1, end)
}

func main() {
	arr := []int{2, 1, 6, 10, 4, 1, 3, 9, 34, 1, 4, 6, 234, 84, 4, 4, 33, 11, 23, 788, 2, 4, 6, 1, 1111, 213123, 41, 5151, 2323}

	QuickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}
