// Package mySqrt :
// File:  mySqrt.go
// Date:  2021/2/7 14:53
// Desc: 第69题：x的平方根
// 计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
package mySqrt

func mySqrt(x int) int {
	left, right, ans := 0, x, -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
