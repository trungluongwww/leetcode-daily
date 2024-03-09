package main

import (
	"fmt"
	"github.com/trungluongwww/leetcode-daily/structs"
	"time"
)

type ListNode = structs.ListNode

func hasCycle(head *ListNode) bool {

	cpy := head

	for cpy.Next != nil {
		fmt.Println(cpy.Val)
		next := cpy.Next
		for next.Next != nil {
			if next.Next == cpy {
				return true
			}
			next = next.Next
		}

		cpy = cpy.Next
	}

	return false
}

func main() {
	s := time.Now()

	fmt.Println(s)
	fmt.Println(s.AddDate(0, 9, 0))
}
