// Package findMin2 :
// File:  findMin2_test.go.go
// Date:  2021/2/7 16:37
// Desc:
package findMin2

import "fmt"

func Example() {
	nums1 := []int{1, 3, 5}
	nums2 := []int{2, 2, 2, 0, 1}
	fmt.Println(findMin(nums1))
	fmt.Println(findMin(nums2))

	// Output:
	// 1
	// 0
}
