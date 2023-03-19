/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 *
 * https://leetcode.cn/problems/house-robber/description/
 *
 * algorithms
 * Medium (54.21%)
 * Likes:    2482
 * Dislikes: 0
 * Total Accepted:    710.9K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,3,1]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[1,2,3,1]
 * 输出：4
 * 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 * 示例 2：
 *
 *
 * 输入：[2,7,9,3,1]
 * 输出：12
 * 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
 * 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */
package main

// @lc code=start
// 动态规划
func rob(nums []int) int {
	length := len(nums)
	// 保护性编程
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	/* dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = max198(nums[0], nums[1])

	for i := 2; i < length; i++ {
		// 转移方程
		dp[i] = max198(dp[i-2]+nums[i], dp[i-1])
	} */

	// 优化空间复杂度
	dp0 := nums[0]
	dp1 := max198(nums[0], nums[1])

	for i := 2; i < length; i++ {
		// 转移方程
		temp := dp1
		dp1 = max198(dp0+nums[i], dp1)
		dp0 = temp
	}

	return dp1
}

func max198(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
