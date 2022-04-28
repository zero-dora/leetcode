package main

import (
	"fmt"
	"math"
)

//76 最小覆盖子串

func minWindow(s string, t string) string {
	result := ""               //返回结果
	subLength := math.MaxInt32 //窗口长度
	i := 0                     //窗口开始位置
	tMap := map[byte]int{}     // t字符串中各个字符的数目
	subMap := map[byte]int{}   //滑动窗口字符串中各个字符的数目
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	//检查窗口字符串是否已经包含了该查询的字符串
	check := func() bool {
		for k, v := range tMap {
			if subMap[k] < v {
				return false
			}
		}
		return true
	}

	for j := 0; j < len(s); j++ {
		if j < len(s) && tMap[s[j]] > 0 { //将t中存在的字符串 添加到窗口的字符map中
			subMap[s[j]]++
		}
		for check() && i <= j { //如果窗口字符串已经全部包含t字符串
			//对比当前窗口长度是否比之前的长度小 如果更小就更换为当前的
			if j-i+1 < subLength {
				subLength = j - i + 1
				result = s[i : j+1]
			}
			//窗口开始位置向前移动 如果i位置的字符串在t里面 则窗口字符map里面的数据需要更新
			if _, ok := subMap[s[i]]; ok {
				subMap[s[i]] -= 1
			}
			i++
		}
	}
	return result
}

func main() {
	s1 := "ADOBECODEBANC"
	t1 := "ABC"
	fmt.Println(minWindow(s1, t1))

	s2 := "a"
	t2 := "a"
	fmt.Println(minWindow(s2, t2))

	s3 := "a"
	t3 := "aa"
	fmt.Println(minWindow(s3, t3))

	s4 := "cabwefgewcwaefgcf"
	t4 := "cae"
	fmt.Println(minWindow(s4, t4))
}
