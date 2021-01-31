// Package longestCommonPrefix :
// File:  longestCommonPrefix_test.go.go
// Date:  2021/1/19 0:00
// Desc:
package longestCommonPrefix

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"2", args{[]string{"dog", "racer", "car"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
