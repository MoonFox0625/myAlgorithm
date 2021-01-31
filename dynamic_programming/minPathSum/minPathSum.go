// Package minPathSum :
// File:  minPathSum.go
// Date:  2021/1/24 15:23
// Desc:  第64题：最小路径和
// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 示例:
//
// 输入:
// [
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
// ]
// 输出: 7
// 解释: 因为路径 1→3→1→1→1 的总和最小。
package minPathSum

func MinPathSum(grid [][]int) int {
	gridLength := len(grid)
	if gridLength < 1 {
		return 0
	}
	dp := make([][]int, gridLength) // dp[i][j] : 表示包含第i行j列元素的最小路径和
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < gridLength; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}
	return dp[gridLength-1][len(dp[gridLength-1])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// OptimizationsMinPathSum : 调用后会修改grid内的内容
func OptimizationsMinPathSum(grid [][]int) int {
	gridLength := len(grid)
	if gridLength < 1 {
		return 0
	}
	// grid[i][j]之后会被修改为 : 表示包含第i行j列元素的最小路径和

	for i := 0; i < gridLength; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j != 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				grid[i][j] = min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
			}
		}
	}
	return grid[gridLength-1][len(grid[gridLength-1])-1]
}
