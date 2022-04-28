package main

import (
	"fmt"
	"math"
)

//209 长度最小的子数组
//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

//暴力解法
func minSubArrayLen1(target int, nums []int) int {
	result := math.MaxInt32 // 最终的结果
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				result = min(result, j-i+1)
				break
			}
		}
	}
	if result == math.MaxInt32 {
		return 0
	}
	return result
}

//滑动窗口
func minSubArrayLen2(target int, nums []int) int {
	result := math.MaxInt32 // 最终的结果
	sum := 0                //滑动窗口数据之和
	i := 0                  //滑动窗口起始位置
	subLength := 0          //滑动窗口长度
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			subLength = j - i + 1
			sum -= nums[i]
			result = min(result, subLength)
			i++
		}
	}

	if result == math.MaxInt32 {
		return 0
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen2(target, nums))

	nums1 := []int{1, 4, 4}
	target1 := 4
	fmt.Println(minSubArrayLen2(target1, nums1))

	nums2 := []int{1, 1, 1, 1, 1, 1, 1, 1}
	target2 := 11
	fmt.Println(minSubArrayLen1(target2, nums2))
}
