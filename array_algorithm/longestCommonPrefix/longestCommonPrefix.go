// Package longestCommonPrefix :
// File:  longestCommonPrefix.go
// Date:  2021/1/18 23:57
// Desc: 题目14: 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
// 所有输入只包含小写字母 a-z

package longestCommonPrefix

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs {
		for strings.Index(str, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
