package main

import "fmt"

// -----------------链表-----------------

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 遍历链表
func (head *ListNode) traverse() {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}

// 反转链表
func (head *ListNode) reverse() *ListNode {
	var preNode *ListNode
	curNode := head
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = nextNode
	}

	return preNode
}

// --------------------------------------

// ---------------二叉树------------------

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --------------------------------------
