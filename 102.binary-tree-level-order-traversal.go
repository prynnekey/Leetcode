/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (65.45%)
 * Likes:    1620
 * Dislikes: 0
 * Total Accepted:    774.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[3],[9,20],[15,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：[[1]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 2000] 内
 * -1000 <= Node.val <= 1000
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
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return nil
	}

	// 队列
	queue := make([]*TreeNode, 0)

	// 根入队
	queue = append(queue, root)

	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})
		n := len(queue)
		for j := 0; j < n; j++ {
			// 出队
			root = queue[0]
			queue = queue[1:]

			// 保存数据
			res[i] = append(res[i], root.Val)

			// 左节点入队
			if root.Left != nil {
				queue = append(queue, root.Left)
			}

			// 右节点入队
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
	}

	return res
}

// @lc code=end
