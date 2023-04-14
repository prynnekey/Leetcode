package main

import "fmt"

// 基数排序
func radixSort(a []int) {
	n := len(a)
	// 找出最大值 确定创建几位桶
	max := getMaxInArray(a)

	// 创建几位桶 个位 十位 百位 ...
	bucketBit := 0

	for max != 0 {
		max /= 10
		bucketBit++
	}

	mod := 10
	div := 1

	for bucketBit > 0 {
		buckets := make([][]int, 10)

		// 将每一位 从小到大加入对应的桶中
		for i := 0; i < n; i++ {
			// 获取位数
			index := (a[i] % mod) / div

			buckets[index] = append(buckets[index], a[i])
		}

		offset := 0
		// 遍历桶 将桶中的数据覆盖到a中
		for i := 0; i < 10; i++ {
			copy(a[offset:], buckets[i])

			offset += len(buckets[i])
		}

		bucketBit--
		mod *= 10
		div *= 10
	}
}

// 获取数组中的最大值
func getMaxInArray(a []int) (max int) {
	max = a[0]

	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	return max
}

func main() {
	a := []int{31, 16, 37, 2, 13, 32, 10, 27, 7, 42, 29, 222, 543, 18, 28, 12, 9}

	radixSort(a)

	fmt.Printf("a: %v\n", a)
}
