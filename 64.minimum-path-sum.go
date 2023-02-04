/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode.cn/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (69.45%)
 * Likes:    1422
 * Dislikes: 0
 * Total Accepted:    448.5K
 * Total Submissions: 645.6K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1
 * 0
 *
 *
 */
package main

import "math"

// @lc code=start
// 使用了额外空间
// func minPathSum(grid [][]int) int {
// 	m := len(grid)
// 	n := len(grid[0])
// 	// 声明
// 	dp := make([][]int, m)
// 	for i := 0; i < m; i++ {
// 		dp[i] = make([]int, n)
// 	}
//
// 	// 状态转移
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			// 四种情况
// 			if i == 0 && j == 0 {
// 				// 在左上角
// 				dp[i][j] = grid[i][j]
// 			} else if i == 0 {
// 				dp[0][j] = grid[0][j] + dp[0][j-1]
// 			} else if j == 0 {
// 				dp[i][0] = grid[i][0] + dp[i-1][0]
// 			} else {
// 				dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
// 			}
// 		}
// 	}
//
// 	return dp[m-1][n-1]
// }

// 优化
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 四种情况
			if i == 0 && j == 0 {
				// 在左上角
				continue
			} else if i == 0 {
				grid[0][j] += grid[0][j-1]
			} else if j == 0 {
				grid[i][0] += grid[i-1][0]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}

	return grid[m-1][n-1]
}

func min(nums ...int) int {
	minNum := math.MaxInt
	for _, num := range nums {
		if minNum > num {
			minNum = num
		}
	}

	return minNum
}

// @lc code=end
