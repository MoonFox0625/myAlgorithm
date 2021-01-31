// Package plusOne :
// File:  plusOne.go
// Date:  2021/1/19 18:26
// Desc: 第66题：加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，这个整数不会以零开头。
package plusOne

func plusOne(digits []int) []int {
	carry := 1 // 进位
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		carry = 0
		if digits[i] > 9 {
			carry = 1
			digits[i] = digits[i] % 10
		}
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
