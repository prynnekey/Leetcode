/*
 * @lc app=leetcode.cn id=870 lang=golang
 *
 * [870] 优势洗牌
 *
 * https://leetcode.cn/problems/advantage-shuffle/description/
 *
 * algorithms
 * Medium (50.24%)
 * Likes:    365
 * Dislikes: 0
 * Total Accepted:    62K
 * Total Submissions: 123.4K
 * Testcase Example:  '[2,7,11,15]\n[1,10,4,11]'
 *
 * 给定两个大小相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i
 * 的数目来描述。
 *
 * 返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
 * 输出：[2,11,7,15]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
 * 输出：[24,32,8,12]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length <= 10^5
 * nums2.length == nums1.length
 * 0 <= nums1[i], nums2[i] <= 10^9
 *
 *
 */
package main

import "sort"

// @lc code=start
func advantageCount(nums1 []int, nums2 []int) []int {
	length := len(nums1)
	index1 := make([]int, length)
	index2 := make([]int, length)

	// 初始化
	for i := 1; i < length; i++ {
		index1[i] = i
		index2[i] = i
	}

	// 排序
	sort.Slice(index1, func(i, j int) bool {
		return nums1[index1[i]] < nums1[index1[j]]
	})

	sort.Slice(index2, func(i, j int) bool {
		return nums2[index2[i]] < nums2[index2[j]]
	})

	res := make([]int, length)

	left, right := 0, length-1
	for i := 0; i < length; i++ {
		if nums1[index1[i]] > nums2[index2[left]] {
			res[index2[left]] = nums1[index1[i]]
			left++
		} else {
			res[index2[right]] = nums1[index1[i]]
			right--
		}
	}

	return res
}

// @lc code=end
