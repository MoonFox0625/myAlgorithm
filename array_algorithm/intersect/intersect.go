// Package array_algorithm :
// File:  intersect.go
// Date:  2021/1/18 22:28
// Desc: 第350题：两个数组的交集
// 给定两个数组，编写一个函数来计算它们的交集。
package intersect

import "sort"

func Intersect(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int) // use int to suit repeat num
	for _, num := range nums1 {
		numMap[num] += 1
	}
	count := 0
	for _, num := range nums2 {
		if numMap[num] > 0 {
			numMap[num]--
			nums2[count] = num
			count++
		}
	}
	return nums2[:count]
}

func IntersectSort(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			nums2[k] = nums2[j]
			k++
			j++
			i++
		}
	}
	return nums2[:k]
}
