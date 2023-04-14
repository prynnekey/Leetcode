package main

import (
	"fmt"
	"math"
)

func bucketSort(a []int) {
	// 桶的数量
	bucketNum := len(a)

	// 获取数组的最大值
	max := getMaxInArray(a)

	// 桶
	buckets := make([][]int, bucketNum)

	// 将数据装入桶中
	for i := 0; i < bucketNum; i++ {
		// 要插入到哪个桶中
		// NOTE:不懂
		index := a[i] * (bucketNum - 1) / max

		buckets[index] = append(buckets[index], a[i])
	}

	// 查看每个桶中的数据
	for i := 0; i < bucketNum; i++ {
		fmt.Printf("%v\n", buckets[i])
	}

	// 桶内排序
	offset := 0
	for i := 0; i < bucketNum; i++ {
		bucketLen := len(buckets[i])

		if bucketLen > 0 {
			sortInBucket(buckets[i])

			copy(a[offset:], buckets[i])

			offset += bucketLen
		}
	}
}

func sortInBucket(bucket []int) {
	insertSort(bucket)
}

// 插入排序
func insertSort(a []int) {
	n := len(a)
	for i := 1; i < n; i++ {
		curValue := a[i]

		insertIndex := i

		// 逆序遍历已排好序的数组
		for insertIndex > 0 {
			// 判断当前元素是否小于
			if curValue < a[insertIndex-1] {
				a[insertIndex] = a[insertIndex-1]

				insertIndex--
			} else {
				break
			}
		}

		// 插入
		a[insertIndex] = curValue
	}
}

func getMaxInArray(a []int) (max int) {
	max = math.MinInt
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	return max
}

func main() {
	a := []int{31, 16, 37, 2, 13, 32, 10, 27, 7, 42, 29, 18, 28, 12, 9}
	bucketSort(a)
	fmt.Println("BucketSort:", a)
}
