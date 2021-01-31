// Package removeNthFromEnd :
// File:  removeNthFromEnd.go
// Date:  2021/1/20 12:31
// Desc: 第19题：删除链表倒数第N个节点
// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
// 示例：
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
// 说明：
// 给定的 n 保证是有效的。
package removeNthFromEnd

import . "myAlgorithm/basic"

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	guard := &ListNode{Next: head}
	pre, cur := guard, guard
	i := 1
	for head != nil {
		if i >= n { // 因为cur 开始就跟head 相隔了一个节点,后续只要移动n-1个
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = cur.Next
	return guard.Next
}
