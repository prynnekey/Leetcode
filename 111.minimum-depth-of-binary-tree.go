/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (51.73%)
 * Likes:    952
 * Dislikes: 0
 * Total Accepted:    523.8K
 * Total Submissions: 1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明：叶子节点是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [2,null,3,null,4,null,5,null,6]
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数的范围在 [0, 10^5] 内
 * -1000
 *
 *
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 广度优先算法
//
// 通过queue实现
func minDepth(root *TreeNode) int {
	// 保护性编程
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	count := make([]int, 0)

	// 入队
	queue = append(queue, root)
	count = append(count, 1)

	for len(queue) > 0 {
		// 出队
		node := queue[0]
		queue = queue[1:]

		depth := count[0]
		count = count[1:]

		if node.Left == nil && node.Right == nil {
			// 找到了叶子节点
			return depth
		}

		if node.Left != nil {
			// 入队
			queue = append(queue, node.Left)
			count = append(count, depth+1)
		}

		if node.Right != nil {
			// 入队
			queue = append(queue, node.Right)
			count = append(count, depth+1)
		}
	}

	return 0
}

// @lc code=end
