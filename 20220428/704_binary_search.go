package main

import "fmt"

// search 二分查找算法
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1 //target 查找的区间 [left,right] 左闭右闭
	for left <= right {           //当left = right 时 区间[left,right] 是有效的 需要再判断一次 所以用 <=
		middle := left + (right-left)/2 //获取中间位置防止溢出  等同于(left+right)/2
		if nums[middle] > target {      //target 在左区间 [left,middle-1]
			right = middle - 1
		} else if nums[middle] < target { //target 在右区间 [middle+1,right]
			left = middle + 1
		} else {
			return middle // 数组中找到目标值，直接返回下标
		}
	}
	return -1 //未找到目标值
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))
	target = 2
	fmt.Println(search(nums, target))
}
