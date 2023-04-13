package main

import "fmt"

// 堆排序
func heapSort(a []int) {
	for i := len(a); i > 1; i-- {
		// 1. 将数组构建成二叉堆(完全二叉树+最大堆)
		buildMaxHeap(a, i)
		// 2. 将根节点(堆首位)和最后一个节点交换位置
		swap(&a[0], &a[i-1])
	}
}

// 构建最大堆
func buildMaxHeap(a []int, len int) {
	// 找到最后一个非叶子节点
	parent := len/2 - 1
	for parent >= 0 {
		adjustHeap(a, parent, len)
		parent--
	}
	// fmt.Println(a[0:len])
}

// 调整a为最大堆
func adjustHeap(a []int, parent, len int) {
	// 最大值的下标
	maxIndex := parent

	// 左右孩子
	leftIndex, rightIndex := 2*parent+1, 2*(parent+1)

	// 找到最大值
	if leftIndex < len && a[leftIndex] > a[maxIndex] {
		maxIndex = leftIndex
	}

	if rightIndex < len && a[rightIndex] > a[maxIndex] {
		maxIndex = rightIndex
	}

	// 将最大值移动到最顶部
	if maxIndex != parent {
		swap(&a[maxIndex], &a[parent])
		adjustHeap(a, maxIndex, len)
	}
}

// 交换a b的值
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := []int{69, 39, 77, 23, 32, 45, 58, 63, 93, 4, 37, 22, 11, 12, 44, 94, 3432, 123, 56, 734, 341, 2334, 444, 5556, 777, 773, 223, 98, 876, 123}

	heapSort(a)

	fmt.Printf("a: %v\n", a)
}
