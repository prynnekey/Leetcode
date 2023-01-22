package main

func missingNumber(nums []int) int {
	set := make(map[int]bool)
	// 初始化map
	for _, v := range nums {
		set[v] = true
	}

	// 遍历map
	for i := 0; i <= len(nums); i++ {
		if ok := set[i]; !ok {
			return i
		}
	}
	return 0
}
