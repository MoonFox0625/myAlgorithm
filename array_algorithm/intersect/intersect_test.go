// Package array_algorithm :
// File:  intersect_test.go.go
// Date:  2021/1/18 22:30
// Desc:
package intersect

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func ExampleIntersect() {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}
	fmt.Println(Intersect(num1, num2))

	num1 = []int{4, 9, 5}
	num2 = []int{9, 4, 9, 8, 4}
	fmt.Println(Intersect(num1, num2))
	// Output:
	// [2 2]
	// [9 4]
}

func TestIntersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2, 2}},
		{"case2", args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.nums1, tt.args.nums2); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestIntersectSort(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2, 2}},
		{"case2", args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectSort(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIntersect(b *testing.B) {
	nums1 := generateRandomSlice(100)
	nums2 := generateRandomSlice(100)
	for i := 0; i < b.N; i++ {
		Intersect(nums1, nums2)
	}
}

func BenchmarkIntersectSort(b *testing.B) {
	nums1 := generateRandomSlice(100)
	nums2 := generateRandomSlice(100)
	for i := 0; i < b.N; i++ {
		IntersectSort(nums1, nums2)
	}
}

// 生成1000以内的随机数切片
func generateRandomSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000)
	}
	return slice
}
