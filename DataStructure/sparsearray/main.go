package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//	1.创建原始的数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 白子

	// 2.输出看原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 3. 转成稀疏数组
	// 不知道数据个数，只能使用切片
	// 思路：
	// （1）遍历chessMap，如果发现有一个元素不为零，创建一个node结构体，
	//	(2) 将其放入对应的切片之中
	var sparseArr []ValNode
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	// 标准的稀疏数组应该含有一个记录原始二维数组的规模（行和列，默认值）
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建ValNode值结点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}

	}

	// 输出稀疏数组
	for i, val := range sparseArr {
		fmt.Printf("%d:%d %d %d\n", i, val.row, val.col, val.val)
	}

	// 将这个稀疏数组，存盘

	// 如何恢复原始的数组

	// 1. 打开文件，恢复原始数组

	// 2. 这里使用稀疏数组恢复

	// 	创建一个原始数组
	var chessMap2 [11][11]int
	//	 遍历稀疏数组，（遍历文件每一行）

	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val

		}

	}

	// 	查看是否恢复
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 存盘，从文件恢复、

	// 改进，将构建的稀疏数组存盘
	// 读盘恢复二维数组
}
