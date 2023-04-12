package main

import "fmt"

// 插入排序
func insertSort(a []int) {
	n := len(a)

	for i := 1; i < n; i++ {
		// 待插入位置
		insertIndex := i

		// 当前元素
		curValue := a[i]

		// 遍历已排序的元素 找到要插入的位置
		for j := i - 1; j >= 0; j-- {
			// 当前元素是否小于已排序的元素
			if curValue < a[j] {
				// 交换位置
				a[insertIndex] = a[j]
				insertIndex--
			} else {
				break
			}
		}

		// 插入
		a[insertIndex] = curValue
	}
}

func main() {
	// a := []int{69, 39, 77, 23, 32, 45, 58, 63, 93, 4, 37, 22}
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	insertSort(a)

	fmt.Printf("a: %v\n", a)
}
