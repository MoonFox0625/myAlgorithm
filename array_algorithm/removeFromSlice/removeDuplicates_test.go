// Package removeFromSlice :
// File:  removeDuplicates_test.go.go
// Date:  2021/1/19 15:25
// Desc:
package removeFromSlice

import (
	"fmt"
	"testing"
)

func ExampleRemoveDuplicates() {
	nums := []int{1, 1, 2}
	fmt.Println(RemoveDuplicates(nums))
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(RemoveDuplicates(nums))
	// Output:
	// 2
	// 5
}

func ExampleRemoveDuplicatesV2() {
	nums := []int{1, 1, 2}
	fmt.Println(RemoveDuplicatesV2(nums))
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(RemoveDuplicatesV2(nums))
	// Output:
	// 2
	// 5
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	for i := 0; i < b.N; i++ {
		RemoveDuplicates(nums)
	}
}

func BenchmarkRemoveDuplicatesV2(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	for i := 0; i < b.N; i++ {
		RemoveDuplicatesV2(nums)
	}
}
