// Package maxProfit :
// File:  maxProfit_test.go.go
// Date:  2021/1/19 11:13
// Desc:
package maxProfit

import "fmt"

func Example() {
	prices := []int{7, 1, 5, 3, 6, 4}
	prices1 := []int{1, 2, 3, 4, 5}
	prices2 := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit(prices1))
	fmt.Println(maxProfit(prices2))
	// Output:
	// 7
	// 4
	// 0
}
