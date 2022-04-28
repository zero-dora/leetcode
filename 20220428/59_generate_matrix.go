package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	startX, startY := 0, 0 //循环起始位置
	loop := n / 2          // 每个圈循环几次，例如n为奇数3，那么loop = 1 只是循环一圈，矩阵中间的值需要单独处理
	mid := n / 2           // 矩阵中间的位置，例如：n为3， 中间的位置就是(1，1)，n为5，中间位置为(2, 2)
	offset := 1            // 每一圈循环，需要控制每一条边遍历的长度
	count := 1             //空格中放置的值
	i, j := 0, 0
	for ; loop > 0; loop-- {
		i = startX
		j = startY
		// 下面开始的四个for就是模拟转了一圈
		// 模拟填充上行从左到右(左闭右开)
		for j = startY; j < startY+n-offset; j++ {
			matrix[i][j] = count
			count++
		}

		//模拟填充右列从上到下(左闭右开)
		for i = startX; i < startX+n-offset; i++ {
			matrix[i][j] = count
			count++
		}
		//模拟填充下行从右到左(左闭右开)
		for ; j > startY; j-- {
			matrix[i][j] = count
			count++
		}

		//模拟填充左列从下到上(左闭右开)
		for ; i > startX; i-- {
			matrix[i][j] = count
			count++
		}

		// 第二圈开始的时候，起始位置要各自加1， 例如：第一圈起始位置是(0, 0)，第二圈起始位置是(1, 1)
		startX++
		startY++
		//offset 控制每一圈里每一条边遍历的长度 2是因为每循环一圈总循环长度会少2
		offset += 2
	}
	// 如果n为奇数的话，需要单独给矩阵最中间的位置赋值
	if (n % 2) == 1 {
		matrix[mid][mid] = count
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
}
