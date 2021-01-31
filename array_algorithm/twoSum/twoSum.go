// Package twoSum :
// File:  twoSum.go
// Date:  2021/1/19 18:57
// Desc: 第1题：两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
package twoSum

func TwoSum(nums []int, target int) []int {
	numIndex := make(map[int]int)
	for i, v := range nums {
		if j, exist := numIndex[target-v]; exist {
			return []int{i, j}
		}
		numIndex[v] = i
	}
	return nil
}
