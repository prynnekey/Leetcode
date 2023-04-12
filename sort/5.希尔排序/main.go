package main

import (
	"fmt"
	"time"
)

// 希尔排序
//
// 为插入排序的变种
func shellSort(a []int) {
	n := len(a)

	// 增量
	gap := n / 2

	for gap > 0 {

		// 默认第一个元素有序
		for i := gap; i < n; i++ {
			// 待插入元素的下标
			insertIndex := i

			// 当前元素
			curValue := a[i]

			// 逆序遍历已排好序的数据
			for insertIndex >= gap {
				if curValue < a[insertIndex-gap] {
					a[insertIndex] = a[insertIndex-gap]
					insertIndex -= gap
				} else {
					break
				}
			}

			// 插入
			a[insertIndex] = curValue
		}

		gap /= 2
	}
}

func main() {
	a := []int{69, 39, 77, 23, 32, 45, 58, 63, 93, 4, 37, 22, 11, 12, 44, 94, 3432, 123, 56, 734, 341, 2334, 444, 5556, 777, 773, 223, 98, 876, 123}
	// a := []int{7, 3, 5, 9, 2}

	start := time.Now()
	shellSort(a)

	elapsed := time.Since(start)

	fmt.Printf("a: %v\n", a)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
