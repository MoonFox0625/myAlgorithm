// Package findMin :
// File:  findMin.go
// Date:  2021/2/7 15:32
// Desc:	第153题：旋转排序数组最小值Ⅰ
// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。请找出其中最小的元素。你可以假设数组中不存在重复元素。
// 示例 1:
//
// 输入: [3,4,5,1,2]
// 输出: 1
// 示例 2:
//
// 输入: [4,5,6,7,0,1,2]
// 输出: 0
package findMin

func findMinOtherOne(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] <= nums[right] {
			return nums[left]
		}
		mid := left + (right-left)/2
		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

// 个人觉得比上面思路好
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
