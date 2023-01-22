package main

/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 *
 * https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (65.88%)
 * Likes:    1140
 * Dislikes: 0
 * Total Accepted:    253.5K
 * Total Submissions: 384.8K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums
 * 中的数字，并以数组的形式返回结果。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,3,2,7,8,2,3,1]
 * 输出：[5,6]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,1]
 * 输出：[2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1
 * 1
 *
 *
 * 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。
 *
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	set := make(map[int]bool)
	res := make([]int, 0)
	// 将所有出现过的数存入set中
	for _, v := range nums {
		set[v] = true
	}

	// 在set中查找[1, n]中没有存在的数
	for i := 1; i < len(nums)+1; i++ {
		ok := set[i]
		if !ok {
			res = append(res, i)
		}
	}
	return res
}

// @lc code=end
