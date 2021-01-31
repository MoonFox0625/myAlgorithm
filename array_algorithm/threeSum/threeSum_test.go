// Package threeSum :
// File:  threeSum_test.go.go
// Date:  2021/1/19 20:59
// Desc:
package threeSum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
