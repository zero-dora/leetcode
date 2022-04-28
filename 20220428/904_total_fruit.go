package main

import (
	"fmt"
)

//滑动窗口
func totalFruit1(fruits []int) int {
	result := 0              // 最终的水果数量
	i := 0                   // 采摘的水果树起始位置
	hashMap := map[int]int{} //水果篮中含有的水果
	for j := 0; j < len(fruits); j++ {
		//直接向水果篮中添加水果
		hashMap[fruits[j]]++
		//当哈希表的长度是3的时候表示区间内元素的种类有3个，已经超过要求

		for len(hashMap) >= 3 {
			hashMap[fruits[i]]--         //需要去除第一个树采摘的水果
			if hashMap[fruits[i]] == 0 { //如果第一颗数对应的水果种类含有的水果数量为0 则在hashMap中删除这种种类 总种类数减一
				delete(hashMap, fruits[i])
			}
			i++
		}
		result = max(result, j-i+1)
	}
	return result
}

//双指针
func totalFruit2(fruits []int) int {
	result := 0                            //最终采摘的水果树数量
	fast, slow := 0, 0                     //可采摘的最后一棵树和第一棵树
	fruitA, fruitB := fruits[0], fruits[0] //当前两个篮子的水果种类
	for fast < len(fruits) {
		if fruits[fast] != fruitA && fruits[fast] != fruitB { //表示采摘的是第三种类的水果
			slow = fast - 1
			fruitA = fruits[slow] //篮子里水果种类调整 a篮水果为fast上一颗树的 肯定和当前不一样
			fruitB = fruits[fast]
			//如果slow大于1 并且slow-1的位置的种类和a篮中种类相同 则slow向前移动 表示前面那颗树也有可能符合采摘条件
			for slow >= 1 && fruits[slow-1] == fruitA {
				slow--
			}
			result = max(result, fast-slow+1)
		} else {
			result = max(result, fast-slow+1)
			fast++
		}
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fruits1 := []int{1, 2, 1}
	fruits2 := []int{0, 1, 2, 2}
	fruits3 := []int{1, 2, 3, 2, 2}
	fruits4 := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println(totalFruit2(fruits1))
	fmt.Println(totalFruit2(fruits2))
	fmt.Println(totalFruit2(fruits3))
	fmt.Println(totalFruit2(fruits4))
}
