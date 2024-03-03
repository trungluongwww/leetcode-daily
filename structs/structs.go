package structs

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitListNode(values []int) *ListNode {
	l := new(ListNode)

	cpy := l

	for i, v := range values {
		cpy.Val = v

		if i != len(values)-1 {
			cpy.Next = new(ListNode)
			cpy = cpy.Next
		}
	}

	return l
}

func (l *ListNode) Print() {
	cpy := l
	for cpy != nil {
		fmt.Print(cpy.Val, "->")
		cpy = cpy.Next
	}
	fmt.Println()
}
