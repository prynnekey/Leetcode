/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode.cn/problems/sqrtx/description/
 *
 * algorithms
 * Easy (38.63%)
 * Likes:    1285
 * Dislikes: 0
 * Total Accepted:    704.6K
 * Total Submissions: 1.8M
 * Testcase Example:  '4'
 *
 * 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
 *
 * 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
 *
 * 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：x = 4
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：x = 8
 * 输出：2
 * 解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= x <= 2^31 - 1
 *
 *
 */
package main

// @lc code=start

// 暴力
// func mySqrt(x int) int {
// 	sqrt := 0
//
// 	for i := 1; i*i <= x; i++ {
// 		sqrt = i
// 	}
//
// 	return sqrt
// }

// 二分查找法
// func mySqrt(x int) int {
// 	res := -1
// 	l, r := 0, x
//
// 	// 二分查找
// 	for l <= r {
// 		mid := l + (r-l)/2
// 		if mid*mid <= x {
// 			// 说明mid小于x的平方根
// 			res = mid
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
//
// 	return res
// }

// 牛顿迭代
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	fx := float64(x)

	return int(sqrt(100, fx))
}

func sqrt(x, i float64) float64 {
	res := (x + i/x) / 2
	if res == x {
		return x
	} else {
		return sqrt(res, i)
	}
}

// @lc code=end
