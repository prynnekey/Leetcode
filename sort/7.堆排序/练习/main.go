package main

import "fmt"

// 堆排序
func heapSort(a []int) {
	for i := len(a); i > 1; i-- {
		// 将数组构建成最大二叉堆 父节点总是大于两个子节点
		buildHeap(a, i)

		// 将堆首和最后一个元素交换位置
		swap(&a[0], &a[i-1])
	}
}

// 构建最大二叉堆
func buildHeap(a []int, len int) {
	// 从第一个非叶子节点开始构建 因为叶子节点没有孩子
	parent := len/2 - 1
	for parent >= 0 {
		// 调整为二叉堆
		adjustHeap(a, parent, len)

		parent--
	}
}

func adjustHeap(a []int, parent, len int) {
	maxIndex := parent

	leftIndex := 2*parent + 1
	rightIndex := 2 * (parent + 1)

	if leftIndex < len && a[leftIndex] > a[maxIndex] {
		maxIndex = leftIndex
	}

	if rightIndex < len && a[rightIndex] > a[maxIndex] {
		maxIndex = rightIndex
	}

	// 交换
	if maxIndex != parent {
		swap(&a[maxIndex], &a[parent])

		// 交换完之后可能孩子就不是二叉堆了
		adjustHeap(a, maxIndex, len)
	}
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := []int{69, 39, 77, 23, 32, 45, 58, 63, 93, 4, 37, 22, 11, 12, 44, 94, 3432, 123, 56, 734, 341, 2334, 444, 5556, 777, 773, 223, 98, 876, 123}

	heapSort(a)

	fmt.Printf("a: %v\n", a)
}
