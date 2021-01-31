// Package reverseString :
// File:  reverseString_test.go.go
// Date:  2021/1/24 22:48
// Desc:
package reverseString

import "fmt"

func Example() {
	str := []byte("Hello")
	ReverseString(str)
	fmt.Println(string(str))
	// Output:
	// olleH
}
