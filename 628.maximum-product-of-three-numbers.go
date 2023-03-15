/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 *
 * https://leetcode.cn/problems/maximum-product-of-three-numbers/description/
 *
 * algorithms
 * Easy (52.21%)
 * Likes:    427
 * Dislikes: 0
 * Total Accepted:    113.7K
 * Total Submissions: 218.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：6
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,4]
 * 输出：24
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [-1,-2,-3]
 * 输出：-6
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3
 * -1000
 *
 *
 */
package main

import (
	"sort"
)

// @lc code=start
func maximumProduct(nums []int) int {
	// 排序
	sort.Ints(nums)

	n := len(nums) - 1

	// 全正+    找出三个最大值
	// 全负-    找出三个最大值
	// nums[n] * nums[n-1] * nums[n-2]

	// 有+ -    ++- --> 两个+最大 -最小
	// nums[n] * nums[n-1] * nums[0]

	//          +-- --> 两个-乘机最大(两个-最小) +最大
	// nums[0] * nums[1] * nums[n]

	return max628(nums[n]*nums[n-1]*nums[n-2], max628(nums[n]*nums[n-1]*nums[0], nums[0]*nums[1]*nums[n]))
}

func max628(a, b int) (max int) {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
