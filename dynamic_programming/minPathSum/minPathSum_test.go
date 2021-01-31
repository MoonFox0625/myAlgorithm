// Package minPathSum :
// File:  minPathSum_test.go.go
// Date:  2021/1/24 15:25
// Desc:
package minPathSum

import "fmt"

func Example() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(MinPathSum(grid))
	fmt.Println(OptimizationsMinPathSum(grid))
	// Output:
	// 7
	// 7
}
