// Package firstUniqChar :
// File:  firstUniqChar_test.go.go
// Date:  2021/1/24 23:01
// Desc:
package firstUniqChar

import "fmt"

func Example() {
	s1 := "leetcode"
	s2 := "loveleetcode"
	fmt.Println(FirstUniqChar(s1))
	fmt.Println(FirstUniqChar(s2))
	// Output:
	// 0
	// 2
}
