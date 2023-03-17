/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 *
 * https://leetcode.cn/problems/number-of-provinces/description/
 *
 * algorithms
 * Medium (62.29%)
 * Likes:    938
 * Dislikes: 0
 * Total Accepted:    259.3K
 * Total Submissions: 416.5K
 * Testcase Example:  '[[1,1,0],[1,1,0],[0,0,1]]'
 *
 *
 *
 * 有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c
 * 间接相连。
 *
 * 省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
 *
 * 给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而
 * isConnected[i][j] = 0 表示二者不直接相连。
 *
 * 返回矩阵中 省份 的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * n == isConnected.length
 * n == isConnected[i].length
 * isConnected[i][j] 为 1 或 0
 * isConnected[i][i] == 1
 * isConnected[i][j] == isConnected[j][i]
 *
 *
 *
 *
 */
package main

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	// 连通域的数量
	count := 0
	// 是否访问过当前节点
	visited := make([]bool, len(isConnected))

	// 深度优先算法
	var dfs func(int)
	dfs = func(i int) {
		// 将当前节点置为访问过
		visited[i] = true

		for j, v := range isConnected[i] {
			if v == 1 && !visited[j] {
				// 连通 且 没有访问过
				dfs(j)
			}

		}

	}

	for i, v := range visited {
		if !v {
			// 还没访问过
			count++

			dfs(i)
		}
	}

	return count
}

// @lc code=end
