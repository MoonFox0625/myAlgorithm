// Package rob :
// File:  rob.go
// Date:  2021/1/24 20:24
// Desc: 第198题：打家劫舍
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
//
// 示例 1:
//
// 输入: [1,2,3,1]
// 输出: 4
// 解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//	 偷窃到的最高金额 = 1 + 3 = 4 。
// 示例 2:
//
// 输入: [2,7,9,3,1]
// 输出: 12
// 解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
package rob

func Rob(nums []int) int {
	numsLength := len(nums)
	if numsLength < 1 {
		return 0
	}
	if numsLength == 1 {
		return nums[0]
	}
	if numsLength == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums)) // dp为偷盗每一间房子时可以获取到的最大金额
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < numsLength; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[numsLength-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func OptimizationsRob(nums []int) int {
	numsLength := len(nums)
	if numsLength < 1 {
		return 0
	}
	if numsLength == 1 {
		return nums[0]
	}
	if numsLength == 2 {
		return max(nums[0], nums[1])
	}
	// nums[i]修改为偷盗每一间房子时可以获取到的最大金额
	nums[1] = max(nums[0], nums[1])
	for i := 2; i < numsLength; i++ {
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[numsLength-1]
}
