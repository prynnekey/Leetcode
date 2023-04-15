package main

import (
	"fmt"
	"math/rand"
)

func sortArray(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

// 排序
func sort(a []int, start, end int) {
	// 参数校验
	if len(a) < 1 || start > end || start < 0 || end > len(a)-1 {
		return
	}

	length := end - start + 1
	// 判断元素个数
	if length < 15 {
		// 插入排序
		insertSort(a, start, end)
	} else {
		// 快速排序

		// 分区指示器
		zoneIndex := partition(a, start, end)

		if zoneIndex > start {
			sort(a, start, zoneIndex-1)
		}

		if zoneIndex < end {
			sort(a, zoneIndex+1, end)
		}
	}
}

func insertSort(a []int, start, end int) {
	for i := start + 1; i <= end; i++ {
		curValue := a[i]

		insertIndex := i

		for insertIndex > 0 {
			if curValue < a[insertIndex-1] {
				a[insertIndex] = a[insertIndex-1]
				insertIndex--
			} else {
				break
			}
		}

		// 找到了待插入元素的位置
		a[insertIndex] = curValue
	}
}

// 分区
func partition(a []int, start, end int) int {
	if start == end {
		return start
	}

	// 基准数
	pivot := rand.Intn(end-start+1) + start

	// 将基准数和最后一位交换位置
	swap(&a[pivot], &a[end])

	// 分区指示器
	// zoneIndex := start - 1

	// 遍历
	// for i := start; i <= end; i++ {
	// 	// 判断当前元素是否小于基准数
	// 	if a[i] <= a[end] {
	// 		zoneIndex++
	//
	// 		if zoneIndex < i {
	// 			swap(&a[i], &a[zoneIndex])
	// 		}
	// 	}
	// }

	for start < end {
		// 修改原来的 nums[j] >= nums[i]，增加交换频率
		for a[start] < a[end] && start < end {
			end--
		}

		if start < end {
			swap(&a[start], &a[end])
			start++
		}

		// 修改原来的 nums[j] >= nums[i]，增加交换频率
		for a[start] < a[end] && start < end {
			start++
		}

		if start < end {
			swap(&a[start], &a[end])
			end--
		}
	}

	return start
	// return zoneIndex
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := []int{31, 16, 37, 2, 13, 32, 10, 27, 7, 42, 29, 222, 543, 18, 28, 12, 9}

	// sortArray(a)
	sortArray(a)

	fmt.Printf("a: %v\n", a)
}
