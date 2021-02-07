// Package findMin :
// File:  findMin_test.go.go
// Date:  2021/2/7 15:33
// Desc:
package findMin

import "fmt"

func Example() {
	nums1 := []int{3, 4, 5, 1, 2}
	nums2 := []int{4, 5, 6, 7, 0, 1, 2, 3}
	fmt.Println(findMin(nums1))
	fmt.Println(findMin(nums2))
	// Output:
	// 1
	// 0
}
