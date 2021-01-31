// Package minimumTotal :
// File:  minimumTotal_test.go.go
// Date:  2021/1/24 11:45
// Desc:
package minimumTotal

import "fmt"

func Example() {
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	fmt.Println(MinimumTotal(triangle))
	fmt.Println(OptimizationMinimumTotal(triangle))

	// Output:
	// 11
	// 11
}
