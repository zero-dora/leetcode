package main

import (
	"fmt"
	"sort"
)
//977.有序数组的平方

//sortedSquares1 先循环数组每个数平方 再排序
func sortedSquares1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}

	sort.Ints(nums)
	return nums
}

//sortedSquares2 双指针 将两端指针中最大值一个个放入新的数组
func sortedSquares2(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1
	newNums := make([]int, n)
	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			newNums[k] = nums[i] * nums[i]
			i++
		} else {
			newNums[k] = nums[j] * nums[j]
			j--
		}
		k--
	}
	return newNums
}

func main() {
	nums1 := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares2(nums1))

	nums2 := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares2(nums2))

}
