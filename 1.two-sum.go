package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// 数据 下标
	hash := make(map[int]int)

	for i, v := range nums {
		if index, ok := hash[target-v]; ok {
			// 存在
			return []int{index, i}
		}

		// 不存在
		hash[v] = i
	}

	return nil
}

// @lc code=end

// 给定一个升序数组,求两数之和=target
//
// 双指针法
func twoSumOrder(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			// 找到了
			return []int{left, right}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
