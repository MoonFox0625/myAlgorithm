// Package lengthOfLIS :
// File:  lengthOfLIS.go
// Date:  2021/1/23 10:27
// Desc:	第300题：最长上升子序列
// 给定一个无序的整数数组，找到其中最长上升子序列的长度。
// 示例:
//
// 输入: [10,9,2,5,3,7,101,18]
// 输出: 4
// 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
// 说明:
//
// 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
package lengthOfLIS

func LengthOFLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make(map[int]int, len(nums)) // dp[i] ：表示以nums[i]结尾的最长上升子序列的长度
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
