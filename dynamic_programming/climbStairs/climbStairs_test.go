// Package climbStairs :
// File:  climbStairs_test.go.go
// Date:  2021/1/20 20:44
// Desc:
package climbStairs

import "testing"

func Test_climbStair(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"2", args{2}, 2},
		{"1", args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStair(tt.args.n); got != tt.want {
				t.Errorf("climbStair() = %v, want %v", got, tt.want)
			}
		})
	}
}
