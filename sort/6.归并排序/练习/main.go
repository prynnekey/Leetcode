package main

import "fmt"

// 归并排序
func mergeSort(a []int) []int {
	n := len(a)
	if n < 2 {
		// 此时n==1
		return a
	}
	mid := n / 2

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
		// 将更小的加入结果集
		if left[leftIndex] < right[rightIndex] {
			res = append(res, left[leftIndex])
			leftIndex++
		} else {
			res = append(res, right[rightIndex])
			rightIndex++
		}
	}

	// 将剩余的数组全部加入res中
	if leftIndex < leftLen {
		res = append(res, left[leftIndex:]...)
	}

	if rightIndex < rightLen {
		res = append(res, right[rightIndex:]...)
	}

	return res
}

func main() {
	a := []int{69, 39, 77, 23, 32, 45, 58, 63, 93, 4, 37, 22, 11, 12, 44, 94, 3432, 123, 56, 734, 341, 2334, 444, 5556, 777, 773, 223, 98, 876, 123}

	res := mergeSort(a)

	fmt.Printf("res: %v\n", res)
}
