package main

import "fmt"
//在排序数组中查找元素的第一个和最后一个位置


//获取右边界 不包含target
func getRightBorder(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	rightBorder := -2 //区别情况一和情况二
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle - 1 // target 在左区间，所以[left, middle - 1]
		} else { // 当nums[middle] == target的时候，更新left，这样才能得到target的右边界
			left = middle + 1
			rightBorder = left
		}
	}
	return rightBorder //注意此时的边界是不包含target值 因为当 nums[middle] == target 并没有直接返回mid而是 在middle上+1了
}

//获取左边界 不包含target
func getLeftBorder(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	leftBorder := -2 //区别情况一和情况二
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] >= target {
			right = middle - 1 // target 在左区间，所以[left, middle - 1]
			// 当nums[middle] == target的时候，更新left，这样才能得到target的左边界 且不含target
			leftBorder = right
		} else {
			left = middle + 1
		}
	}
	return leftBorder //注意此时的边界是不包含target值 因为当 nums[middle] == target 并没有直接返回mid而是 在middle上-1了
}

func searchRange(nums []int, target int) []int {
	leftBorder := getLeftBorder(nums, target)
	rightBorder := getRightBorder(nums, target)

	//情况一
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}
	//情况三
	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}
	//情况二
	return []int{-1, -1}
}

//getBorder 获取边界值 isLeft true 获取左边界 false 获取右边界
//需要左边界则右指针一直往前移动（遇到相等情况）,右边界则左指针一直往右移动（遇到相等）
func getBorder(nums []int, target int, isLeft bool) int {
	left := 0
	right := len(nums) - 1
	ans := len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target || (isLeft && nums[mid] >= target) { //查询左边界
			right = mid - 1 // target 在左区间，所以[left, middle - 1]
			// 当nums[middle] == target的时候，更新left，这样才能得到target的左边界 且不含target
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans

}

func searchRange2(nums []int, target int) []int {
	leftBorder := getBorder(nums, target, true)
	rightBorder := getBorder(nums, target, false) - 1
	if leftBorder <= rightBorder && rightBorder < len(nums) && nums[leftBorder] == target && nums[rightBorder] == target {
		return []int{leftBorder, rightBorder}
	}
	return []int{-1, -1}
}

//searchTarget 查找第一个大于或等于target的位置
func searchTarget(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func searchRange3(nums []int, target int) []int {
	leftBorder := searchTarget(nums, target)

	if leftBorder == len(nums) || nums[leftBorder] != target {
		return []int{-1, -1}
	}
	rightBorder := searchTarget(nums, target+1) - 1
	return []int{leftBorder, rightBorder}
}

//情况一：target 在数组范围的右边或者左边，例如数组{3, 4, 5}，target为2或者数组{3, 4, 5},target为6，此时应该返回{-1, -1}
//情况二：target 在数组范围中，且数组中不存在target，例如数组{3,6,7},target为5，此时应该返回{-1, -1}
//情况三：target 在数组范围中，且数组中存在target，例如数组{3,6,7},target为6，此时应该返回{1, 1}
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange2(nums, target))
	target = 6
	fmt.Println(searchRange2(nums, target))
	target = 0
	fmt.Println(searchRange2(nums, target))
}
