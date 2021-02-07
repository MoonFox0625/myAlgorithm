// Package firstBadVersion :
// File:  firstBadVersion.go
// Date:  2021/2/7 15:18
// Desc:   278题：第一个错误的版本
// 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
//
// 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
//
//
// 示例:
//
// 给定 n = 5，并且 version = 4 是第一个错误的版本。
//
// 调用 isBadVersion(3) -> false
// 调用 isBadVersion(5) -> true
// 调用 isBadVersion(4) -> true
//
// 所以，4 是第一个错误的版本。
package firstBadVersion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left

}

func isBadVersion(version int) bool {
	// leetcode 自己提供
	// Not implementation;
	return false
}
