package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 获取链表的长度len
	p := head
	len := 0
	for p != nil {
		p = p.Next
		len++
	}

	// 倒数第k个节点相当于len-k+1个节点
	for i := 0; i < len-k; i++ {
		head = head.Next
	}

	return head
}
