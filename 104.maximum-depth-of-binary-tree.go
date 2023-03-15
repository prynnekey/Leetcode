/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (77.14%)
 * Likes:    1542
 * Dislikes: 0
 * Total Accepted:    962.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 *
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最大深度 3 。
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

// 深度优先算法
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
//
// 	// 找到了叶子节点
// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}
//
// 	max := math.MinInt
//
// 	if root.Left != nil {
// 		max = max104(max, maxDepth(root.Left))
// 	}
//
// 	if root.Right != nil {
// 		max = max104(max, maxDepth(root.Right))
// 	}
//
// 	return max + 1
// }

// 广度优先算法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 队列
	queue := make([]*TreeNode, 0)
	// 深度
	count := make([]int, 0)

	// 初始化
	queue = append(queue, root)
	count = append(count, 1)

	var depth int

	for len(queue) > 0 {
		// 出队
		node := queue[0]
		queue = queue[1:]

		depth = count[0]
		count = count[1:]

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

	return depth
}

func max104(a, b int) (max int) {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
