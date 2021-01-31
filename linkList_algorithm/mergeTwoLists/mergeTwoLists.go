// Package mergeTwoLists :
// File:  mergeTwoLists.go
// Date:  2021/1/20 12:59
// Desc: 第21题：合并两个有序链表
// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 示例：
//
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
package mergeTwoLists

import . "myAlgorithm/basic"

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	guard := &ListNode{}
	preHead := guard
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			guard.Next = l1
			l1 = l1.Next
		} else {
			guard.Next = l2
			l2 = l2.Next
		}
		guard = guard.Next
	}
	if l1 != nil {
		guard.Next = l1
	}
	if l2 != nil {
		guard.Next = l2
	}
	return preHead.Next
}
