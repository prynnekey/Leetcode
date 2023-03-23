package main

import "fmt"

// 选择排序
func selectionSort(array []int) {
	n := len(array)
	// 每次找到最小的元素，并移动到数组的最前面
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if array[min] > array[j] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
}

func main() {
	a := []int{10, 2, 8, 6}

	selectionSort(a)

	fmt.Printf("a: %v\n", a)
}
