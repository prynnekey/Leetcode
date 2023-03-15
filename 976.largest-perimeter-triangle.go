/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长
 *
 * https://leetcode.cn/problems/largest-perimeter-triangle/description/
 *
 * algorithms
 * Easy (57.65%)
 * Likes:    233
 * Dislikes: 0
 * Total Accepted:    81.9K
 * Total Submissions: 142.1K
 * Testcase Example:  '[2,1,2]'
 *
 * 给定由一些正数（代表长度）组成的数组 nums ，返回 由其中三个长度组成的、面积不为零的三角形的最大周长 。如果不能形成任何面积不为零的三角形，返回
 * 0。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,1,2]
 * 输出：5
 * 解释：你可以用三个边长组成一个三角形:1 2 2。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,1,10]
 * 输出：0
 * 解释：
 * 你不能用边长 1,1,2 来组成三角形。
 * 不能用边长 1,1,10 来构成三角形。
 * 不能用边长 1、2 和 10 来构成三角形。
 * 因为我们不能用任何三条边长来构成一个非零面积的三角形，所以我们返回 0。
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 10^4
 * 1 <= nums[i] <= 10^6
 *
 *
 */
package main

import (
	"sort"
)

// @lc code=start
func largestPerimeter(nums []int) int {
	// 排序
	sort.Ints(nums)

	// 从后往前取三个数，满足三角形定义即可
	for i := len(nums) - 1; i >= 2; i-- {
		c := nums[i]
		b := nums[i-1]
		a := nums[i-2]

		// 任意两边之和大于第三边 ---> 最小的两个和 > 最大的
		if a+b > c {
			// 成立

			// 任意两边之差小于第三边
			return a + b + c
		}
	}

	return 0
}

// @lc code=end
