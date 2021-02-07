// Package minEatingSpeed :
// File:  minEatingSpeed_test.go.go
// Date:  2021/2/7 1:14
// Desc:
package minEatingSpeed

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		H     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 6, 7, 11}, 8}, 4},
		{"2", args{[]int{30, 11, 23, 4, 20}, 5}, 30},
		{"3", args{[]int{30, 11, 23, 4, 20}, 6}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinEatingSpeed(tt.args.piles, tt.args.H); got != tt.want {
				t.Errorf("MinEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
