package main

import "fmt"

//removeElement 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return j
}

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)

	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums1, 2))
	fmt.Println(nums1)
}
