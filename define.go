package main

import "fmt"

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 遍历链表
func (head *ListNode) traverse() {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
