/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 *
 * https://leetcode.cn/problems/maximum-subarray/description/
 *
 * algorithms
 * Medium (54.86%)
 * Likes:    5626
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 2.3M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 子数组 是数组中的一个连续部分。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
 * 输出：6
 * 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [5,4,-1,7,8]
 * 输出：23
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 *
 * 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
 *
 */
package main

import "math"

// @lc code=start

func maxSubArray(nums []int) int {
	// 前缀和
	sum := make([]int, len(nums)+1)
	for i, num := range nums {
		sum[i+1] = sum[i] + num
	}

	// 前缀和中的最小值
	preMin := make([]int, len(sum))
	for i := 1; i < len(sum); i++ {
		preMin[i] = min(preMin[i-1], sum[i])
	}

	res := math.MinInt

	for i := 1; i < len(sum); i++ {
		res = max53(res, sum[i]-preMin[i-1])
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max53(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// max(s[i] - s[j]) (j: 0 ~ i-1) --> 只需s[j]最小 ---> 求前缀和中的最小值preMin[i] ---> max(s[i] - preMin[i-1])
//

// @lc code=end
