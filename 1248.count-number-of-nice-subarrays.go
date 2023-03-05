/*
 * @lc app=leetcode.cn id=1248 lang=golang
 *
 * [1248] 统计「优美子数组」
 *
 * https://leetcode.cn/problems/count-number-of-nice-subarrays/description/
 *
 * algorithms
 * Medium (57.79%)
 * Likes:    251
 * Dislikes: 0
 * Total Accepted:    48.7K
 * Total Submissions: 84.2K
 * Testcase Example:  '[1,1,2,1,1]\n3'
 *
 * 给你一个整数数组 nums 和一个整数 k。如果某个连续子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
 *
 * 请返回这个数组中 「优美子数组」 的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2,1,1], k = 3
 * 输出：2
 * 解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,4,6], k = 1
 * 输出：0
 * 解释：数列中不包含任何奇数，所以不存在优美子数组。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
 * 输出：16
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 50000
 * 1 <= nums[i] <= 10^5
 * 1 <= k <= nums.length
 *
 *
 */
package main

// @lc code=start
// 连续子数组 --> 前缀和
func numberOfSubarrays(nums []int, k int) int {
	res := 0
	// 前缀和
	sum := make([]int, len(nums)+1)
	for i, num := range nums {
		sum[i+1] = sum[i] + num%2
	}

	hash := make(map[int]int)
	for _, s := range sum {
		res += hash[s-k]
		hash[s]++
	}

	return res
}

// 输入：nums = [1,1,2,1,1], k = 3
// 输出：2
// 解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。

// 1 1 2 1 1          3
// %2
// 1 1 0 1 1 --> 前缀和
// 0 1 2 2 3 4 --> s[i]
//
// count
// i 0 1 2 3 4
// v 1 1 2 1 1

// @lc code=end
