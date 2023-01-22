package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		pos, ok := m[target-v]
		if !ok {
			m[v] = i
			continue
		}
		return []int{pos, i}
	}
	return nil
}

// @lc code=end
