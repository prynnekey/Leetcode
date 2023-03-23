package main

import "fmt"

// 冒泡排序
func bubbleSort(array []int) {
	n := len(array)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			// 将大的移动到最后
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func main() {
	a := []int{10, 2, 8, 6}

	bubbleSort(a)

	fmt.Printf("a: %v\n", a)
}
