package main

import "fmt"

// 插入排序
func insertionSort(a []int) {
	length := len(a)
	// i 表示已经排序过的最后一个位置的下标
	for i := 0; i < length-1; i++ {
		// 当前待排序元素
		curValue := a[i+1]

		// 已被排序过的索引
		preIndex := i
		for preIndex >= 0 {
			if curValue < a[preIndex] {
				// 交换位置
				a[preIndex+1] = a[preIndex]
				preIndex--
			} else {
				break
			}
		}

		// 找到了要插入的位置
		a[preIndex+1] = curValue
	}

}

func main() {
	a := []int{69, 39, 77, 23, 32, 45, 58, 63, 93, 4, 37, 22}

	insertionSort(a)

	fmt.Printf("a: %v\n", a)
}
