// Package hasCycle :
// File:  hasCycle.go
// Date:  2021/1/20 13:19
// Desc: 第141题：环形链表
// 给定一个链表，判断链表中是否有环。为了表示给定链表中的环，
// 我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
// 如果 pos 是 -1，则在该链表中没有环。
// 具体见链接: https://leetcode-cn.com/problems/linked-list-cycle/
package hasCycle

import . "myAlgorithm/basic"

// 通过hash表来检测节点之前是否被访问过，来判断链表是否成环。
func HashHasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}

// 双指针解法
func DoublePointHasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next

	}

	return true
}
