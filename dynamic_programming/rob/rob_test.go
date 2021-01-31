// Package rob :
// File:  rob_test.go.go
// Date:  2021/1/24 20:26
// Desc:
package rob

import "fmt"

func Example() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{2, 7, 9, 3, 1}
	fmt.Println(Rob(nums1))
	fmt.Println(OptimizationsRob(nums1))
	fmt.Println(Rob(nums2))
	fmt.Println(OptimizationsRob(nums2))
	// Output:
	// 4
	// 4
	// 12
	// 12
}
