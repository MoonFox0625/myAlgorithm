// Package removeFromSlice :
// File:  removeElement_test.go.go
// Date:  2021/1/19 14:45
// Desc:
package removeFromSlice

import (
	"fmt"
	"math/rand"
	"myAlgorithm/basic"
	"testing"
)

func ExampleRemoveElement() {
	nums, val := []int{3, 2, 2, 3}, 3
	fmt.Println(RemoveElement(nums, val))
	// Output:
	// 2
}

func ExampleRemoveElementV2() {
	nums, val := []int{3, 2, 2, 3}, 3
	fmt.Println(RemoveElementV2(nums, val))
	// Output:
	// 2
}

func BenchmarkRemoveElement(b *testing.B) {
	nums, val := basic.GenerateRandomSlice(100000, 1000), rand.Intn(1000)
	for i := 0; i < b.N; i++ {
		RemoveElement(nums, val)

	}
}
func BenchmarkRemoveElementV2(b *testing.B) {
	nums, val := basic.GenerateRandomSlice(100000, 1000), rand.Intn(1000)
	for i := 0; i < b.N; i++ {
		RemoveElementV2(nums, val)

	}
}
