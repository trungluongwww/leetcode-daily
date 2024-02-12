package main

import (
	"fmt"
)

//func palindromic(s string, left, right int) int {
//	count := 0
//	for left >= 0 && right < len(s) && s[left] == s[right] {
//		left--
//		right++
//		count++
//	}
//
//	return count
//}
//
//func countSubstrings(s string) int {
//	rs := 0
//	for i := range s {
//		rs += palindromic(s, i, i) + palindromic(s, i, i+i)
//		fmt.Println(i, rs)
//	}
//
//	return rs
//}

func countSubstrings(s string) int {
	count := 0
	for i := range s {
		count += isPalindromeSubstrings(s, i, i) + isPalindromeSubstrings(s, i, i+1)
	}
	return count
}

func isPalindromeSubstrings(s string, l, r int) int {
	res := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		res++
		l--
		r++
	}
	return res
}

func main() {
	fmt.Println(countSubstrings("abc"))
}
