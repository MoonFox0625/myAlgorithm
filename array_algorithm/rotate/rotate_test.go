// Package rotate :
// File:  rotate_test.go.go
// Date:  2021/1/19 13:14
// Desc:
package rotate

import (
	"fmt"
	"reflect"
	"testing"
)

func Example() {
	nums1, k1 := []int{1, 2, 3, 4, 5, 6, 7}, 3
	fmt.Println(rotate(nums1, k1))
	nums2, k2 := []int{-1, -100, 3, 99}, 2
	fmt.Println(rotate(nums2, k2))
	// Output:
	// [5 6 7 1 2 3 4]
	// [3 99 -1 -100]
}

func Test_reverse(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{1, 2, 3}}, []int{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
