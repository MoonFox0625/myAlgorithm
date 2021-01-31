// Package minimumTotal :
// File:  minimumTotal.go
// Date:  2021/1/24 11:13
// Desc: 第120题：三角形最小路径和
// 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
// 例如，给定三角形：
//
// [
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
// ]
// 则自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
package minimumTotal

func MinimumTotal(triangle [][]int) int {
	triangleLength := len(triangle)
	if triangleLength < 1 {
		return 0
	}
	if triangleLength == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, triangleLength) // dp[i][j] : 表示包含第i行j列元素的最小路径和
	for i, arr := range triangle {
		dp[i] = make([]int, len(arr))
	}
	dp[0][0] = triangle[0][0]
	dp[1][0] = dp[0][0] + triangle[1][0]
	dp[1][1] = dp[0][0] + triangle[1][1]
	result := 1<<31 - 1 // math.MaxInt32
	for i := 2; i < triangleLength; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}

		}
	}
	for _, v := range dp[triangleLength-1] {
		result = min(result, v)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func OptimizationMinimumTotal(triangle [][]int) int {
	triangleLength := len(triangle)
	if triangleLength < 1 {
		return 0
	}
	if triangleLength == 1 {
		return triangle[0][0]
	}
	// triangle[i][j] : 表示用包含第i行j列元素的最小路径和替代原来存有的数据
	triangle[1][0] = triangle[0][0] + triangle[1][0]
	triangle[1][1] = triangle[0][0] + triangle[1][1]
	result := 1<<31 - 1 // math.MaxInt32
	for i := 2; i < triangleLength; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
			} else {
				triangle[i][j] = min(triangle[i-1][j-1], triangle[i-1][j]) + triangle[i][j]
			}

		}
	}
	for _, v := range triangle[triangleLength-1] {
		result = min(result, v)
	}
	return result
}
