// Package maxSubArray :
// File:  maxSubArray.go
// Date:  2021/1/20 20:50
// Desc: 第53题：最大子序和
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 示例:
//
// 输入: [-2,1,-3,4,-1,2,1,-5,4],
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
package maxSubArray

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make(map[int]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = nums[i] + dp[i-1]
		}

	}
	var result = -1 << 32
	for _, num := range dp {
		result = max(num, result)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func SimplifySubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make(map[int]int, len(nums))
	dp[0] = nums[0]
	var result = dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(result, dp[i])
	}
	return result
}
