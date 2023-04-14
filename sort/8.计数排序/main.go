package main

import (
	"fmt"
	"math"
)

// 计数排序
func countingSort(a []int) {
	n := len(a)
	// 找出数组中的最大值和最小值
	min := math.MaxInt
	max := math.MinInt

	for i := 0; i < n; i++ {
		if a[i] > max {
			max = a[i]
		}

		if a[i] < min {
			min = a[i]
		}
	}

	// 创建max-min+1个长度的计数数组
	countArray := make([]int, max-min+1)

	// 偏移量，用来计算当前数据在计数数组中的位置i
	offset := -min

	// 遍历数组，并将每个元素出现的次数加入到计数数组中
	for i := 0; i < n; i++ {
		countArray[offset+a[i]]++
	}

	// 遍历计数数组，将计数数组的数据放入数组中
	i, j := 0, 0
	for i < n {
		if countArray[j] != 0 {
			a[i] = j - offset

			countArray[j]--
			i++
		} else {
			j++
		}
	}
}

func main() {
	a := []int{69, 39, 77, 23, 32, 45, 58, 63, 93, 4, 37, 22, 11, 12, 44, 94, 3432, 123, 56, 734, 341, 2334, 444, 5556, 777, 773, 223, 98, 876, 123}
	// a := []int{5, 4, 3, 2, 5, 3, 2, 3, 2}

	countingSort(a)

	fmt.Printf("a: %v\n", a)
}
