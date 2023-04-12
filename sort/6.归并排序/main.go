package main

import "fmt"

// 归并排序
//
// 通过递归和分治法实现
func mergeSort(a []int) []int {
	len := len(a)
	if len < 2 {
		// 此时len==1
		return a
	}
	mid := len / 2

	left := mergeSort(a[:mid])
	right := mergeSort(a[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	res := make([]int, 0)

	leftLen := len(left)
	rightLen := len(right)

	leftIndex, rightIndex := 0, 0
	for leftIndex < leftLen && rightIndex < rightLen {
		// 将小的放入结果集
		if left[leftIndex] < right[rightIndex] {
			res = append(res, left[leftIndex])
			leftIndex++
		} else {
			res = append(res, right[rightIndex])
			rightIndex++
		}
	}

	// 如果其中一个数组还有剩余 则全部加入res中
	if leftIndex < leftLen {
		// 说明left数组中还有数据
		res = append(res, left[leftIndex:]...)
	}

	if rightIndex < rightLen {
		// 说明left数组中还有数据
		res = append(res, right[rightIndex:]...)
	}

	return res
}

func main() {
	a := []int{69, 39, 77, 23, 32, 45, 58, 63, 93, 4, 37, 22, 11, 12, 44, 94, 3432, 123, 56, 734, 341, 2334, 444, 5556, 777, 773, 223, 98, 876, 123}

	res := mergeSort(a)

	fmt.Printf("res: %v\n", res)
}
