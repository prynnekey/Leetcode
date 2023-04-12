package main

import "fmt"

// 快速排序
func quickSort(a []int) {
	sort(a, 0, len(a)-1)
}

func sort(a []int, start, end int) {
	// 参数校验
	if len(a) < 1 || start < 0 || end > len(a)-1 || start > end {
		return
	}

	// 获取分区指示器
	zoneIndex := partition(a, start, end)

	// 递归调用sort
	if zoneIndex > start {
		sort(a, start, zoneIndex-1)
	}

	if zoneIndex < end {
		sort(a, zoneIndex+1, end)
	}
}

// 获取分区指示器
func partition(a []int, start, end int) int {
	// 参数校验
	if start == end {
		return start
	}

	// 随机选择基准数
	pivot := start

	// 将基准数和最后一位交换位置
	swap(a, pivot, end)

	// 分区指示器 第一次为-1
	zoneIndex := start - 1

	for i := start; i <= end; i++ {
		if a[i] <= a[end] {
			// 将分区指示器++
			zoneIndex++

			if zoneIndex < i {
				swap(a, zoneIndex, i)
			}
		}
	}

	return zoneIndex
}

// 交换数组中两个下标的元素位置
func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	a := []int{69, 39, 77, 23, 32, 45, 58, 63, 93, 4, 37, 22}
	// a := []int{7, 3, 5, 9, 2}

	quickSort(a)

	fmt.Printf("a: %v\n", a)

}
