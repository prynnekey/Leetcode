/*
 * @lc app=leetcode.cn id=486 lang=golang
 *
 * [486] 预测赢家
 *
 * https://leetcode.cn/problems/predict-the-winner/description/
 *
 * algorithms
 * Medium (59.48%)
 * Likes:    612
 * Dislikes: 0
 * Total Accepted:    66.1K
 * Total Submissions: 111.1K
 * Testcase Example:  '[1,5,2]'
 *
 * 给你一个整数数组 nums 。玩家 1 和玩家 2 基于这个数组设计了一个游戏。
 *
 * 玩家 1 和玩家 2 轮流进行自己的回合，玩家 1 先手。开始时，两个玩家的初始分值都是 0
 * 。每一回合，玩家从数组的任意一端取一个数字（即，nums[0] 或 nums[nums.length - 1]），取到的数字将会从数组中移除（数组长度减
 * 1 ）。玩家选中的数字将会加到他的得分上。当数组中没有剩余数字可取时，游戏结束。
 *
 * 如果玩家 1 能成为赢家，返回 true 。如果两个玩家得分相等，同样认为玩家 1 是游戏的赢家，也返回 true
 * 。你可以假设每个玩家的玩法都会使他的分数最大化。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,5,2]
 * 输出：false
 * 解释：一开始，玩家 1 可以从 1 和 2 中进行选择。
 * 如果他选择 2（或者 1 ），那么玩家 2 可以从 1（或者 2 ）和 5 中进行选择。如果玩家 2 选择了 5 ，那么玩家 1 则只剩下 1（或者 2
 * ）可选。
 * 所以，玩家 1 的最终分数为 1 + 2 = 3，而玩家 2 为 5 。
 * 因此，玩家 1 永远不会成为赢家，返回 false 。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,5,233,7]
 * 输出：true
 * 解释：玩家 1 一开始选择 1 。然后玩家 2 必须从 5 和 7 中进行选择。无论玩家 2 选择了哪个，玩家 1 都可以选择 233 。
 * 最终，玩家 1（234 分）比玩家 2（12 分）获得更多的分数，所以返回 true，表示玩家 1 可以成为赢家。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 20
 * 0 <= nums[i] <= 10^7
 *
 *
 */
package main

// @lc code=start

// nums = [1,5,233,7]
//
// 玩家一: 选1  {5，233}  233 ---> 1+233
//         选7  {1,5}     5   ---> 7+5
//
// 玩家二: {5,7}    7    5 ---> 5+7
//         {1.233}  233  1 ---> 233+1
//
// 存在递归 每次选择 都依赖于下次已选择
//
// 终止条件
//
// 1. l==r  ---> 只有一个元素了  --->  只能选择nums[l]
//
// 2. r-l == 1  --->  只有两个个元素了 ----> 选择更大的max(nums[r],nums[l])
func PredictTheWinner(nums []int) bool {

	var maxScore func(int, int) int
	maxScore = func(left, right int) int {
		// 终止条件
		if left == right {
			// 只有一个元素了
			return nums[left]
		}

		if right-left == 1 {
			// 只有两个元素了,选择比较大的
			return max486(nums[right], nums[left])
		}

		// 先选左边,对手一定给你留较小的值
		leftScore := nums[left] + min486(maxScore(left+2, right), maxScore(left+1, right-1))

		// 先选右边
		rightScore := nums[right] + min486(maxScore(left+1, right-1), maxScore(left, right-2))

		return max486(leftScore, rightScore)
	}

	// 总分数
	score := 0
	for _, v := range nums {
		score += v
	}

	// 玩家一的分数
	p1 := maxScore(0, len(nums)-1)

	// 玩家二的分数
	// score - p1

	return p1 >= score-p1
}

func max486(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min486(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end
