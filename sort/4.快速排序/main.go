package main

import (
	"fmt"
	"math/rand"
)

// 快速排序
func quickSort(a []int) {
	sort(a, 0, len(a)-1)
}

func sort(a []int, start, end int) {
	// 参数校验
	if len(a) < 1 || start > end || start < 0 || end >= len(a) {
		return
	}

	// 分区指示器
	zoneIndex := partition(a, start, end)

	// 排序分区指示器的左边
	if zoneIndex > start {
		sort(a, start, zoneIndex-1)
	}

	// 排序分区指示器的右边
	if zoneIndex < end {
		sort(a, zoneIndex+1, end)
	}
}

// 分区
func partition(a []int, start, end int) int {
	// 参数校验
	if start == end {
		// 无需分区
		return start
	}

	// 基准数
	// rand.Seed(time.Now().UnixNano())
	// pivot := rand.Intn(end - start + 1)
	pivot := rand.Intn(end-start+1) + start

	// 将基准数和最后一位交换位置
	swap(a, end, pivot)

	// 分区指示器  第一次分区为-1
	zoneIndex := start - 1

	for i := start; i <= end; i++ {
		// 和基准数对比
		// 如果小于等于则分区指示器+1
		if a[i] <= a[end] {
			zoneIndex++

			// 如果分区指示器<i则交换a[i]和a[zoneIndex]
			if zoneIndex < i {
				swap(a, i, zoneIndex)
			}
		}
	}

	return zoneIndex
}

// 交换两个元素
func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	a := []int{69, 39, 77, 23, 32, 45, 58, 63, 93, 4, 37, 22}
	// a := []int{7, 3, 5, 9, 2}

	quickSort(a)

	fmt.Printf("a: %v\n", a)
}
