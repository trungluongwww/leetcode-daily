package main

import (
	"github.com/trungluongwww/leetcode-daily/structs"
)

type ListNode = structs.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cpy := head
	cur := head
	var curCpy *ListNode

	for cpy.Next != nil {
		if n > 1 {
			n--
		} else {
			curCpy = cur
			cur = cur.Next
		}
		cpy = cpy.Next
	}

	if cur.Next != nil {
		cur.Val = cur.Next.Val
		cur.Next = cur.Next.Next
	} else if curCpy != nil {
		curCpy.Next = nil
	} else {
		return curCpy
	}

	return head
}

func main() {
	head := structs.InitListNode([]int{1})
	head.Print()
	head = removeNthFromEnd(head, 1)
	head.Print()
}
