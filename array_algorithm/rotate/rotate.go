// Package rotate :
// File:  rotate.go
// Date:  2021/1/19 13:12
// Desc:  题目189: 旋转数组
// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 要求使用空间复杂度为 O(1) 的 原地 算法。
package rotate

func rotate(nums []int, k int) []int {
	reverse(nums)
	reverse(nums[:k%len(nums)]) // 为了更强的可适应性
	reverse(nums[k%len(nums):])
	return nums
}

func reverse(nums []int) []int {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	return nums
}
