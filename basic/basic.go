// Package basic :
// File:  basic.go
// Date:  2021/1/19 15:06
// Desc:
package basic

import (
	"math/rand"
	"time"
)

// 生成n以内的含有size个数的随机正整数切片
func GenerateRandomSlice(size int, n int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(n)
	}
	return slice
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
