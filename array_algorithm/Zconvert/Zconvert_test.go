// Package Zconvert :
// File:  Zconvert_test.go.go
// Date:  2021/1/19 22:05
// Desc:
package Zconvert

import "testing"

func TestZconvert(t *testing.T) {
	type args struct {
		str     string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"2", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zconvert(tt.args.str, tt.args.numRows); got != tt.want {
				t.Errorf("Zconvert() = %v, want %v", got, tt.want)
			}
		})
	}
}
