// Package maxSubArray :
// File:  maxSubArray_test.go.go
// Date:  2021/1/20 20:50
// Desc:
package maxSubArray

import "fmt"

func Example() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(MaxSubArray(nums))
	// Output:
	// 6
}
