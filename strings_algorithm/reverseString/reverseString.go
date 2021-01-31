// Package reverseString :
// File:  reverseString.go
// Date:  2021/1/24 22:43
// Desc:	第344题：反转字符串
// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
//
// 示例 1：
//
// 输入：["h","e","l","l","o"]
// 输出：["o","l","l","e","h"]
// 示例 2：
//
// 输入：["H","a","n","n","a","h"]
// 输出：["h","a","n","n","a","H"]
package reverseString

func ReverseString(str []byte) {
	left, right := 0, len(str)-1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
}
