/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 *
 * https://leetcode.cn/problems/maximum-average-subarray-i/description/
 *
 * algorithms
 * Easy (43.85%)
 * Likes:    283
 * Dislikes: 0
 * Total Accepted:    98K
 * Total Submissions: 223.9K
 * Testcase Example:  '[1,12,-5,-6,50,3]\n4'
 *
 * 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
 *
 * 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
 *
 * 任何误差小于 10^-5 的答案都将被视为正确答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,12,-5,-6,50,3], k = 4
 * 输出：12.75
 * 解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5], k = 1
 * 输出：5.00000
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= k <= n <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */
package main

import (
	"math"
)

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	// 前缀和
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	max := math.MinInt

	// 连续子数组的最大值
	for i, j := 0, k; j < len(preSum); i, j = i+1, j+1 {
		sum := preSum[j] - preSum[i]
		max = max643(max, sum)
	}

	return float64(max) / float64(k)
}

func max643(a, b int) (max int) {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
