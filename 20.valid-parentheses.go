/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode.cn/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (44.38%)
 * Likes:    3730
 * Dislikes: 0
 * Total Accepted:    1.4M
 * Total Submissions: 3.1M
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 每个右括号都有一个对应的相同类型的左括号。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "()"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "()[]{}"
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "(]"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 仅由括号 '()[]{}' 组成
 *
 *
 */
package main

// @lc code=start
// 通过栈来解决
func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			// 入栈
			stack = append(stack, ')')
		case '[':
			// 入栈
			stack = append(stack, ']')
		case '{':
			// 入栈
			stack = append(stack, '}')
		default:
			// 出栈
			if len(stack) == 0 {
				// 说明栈为空 括号的长度不匹配
				// "({}"
				return false
			}

			// 可以出栈
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if pop != s[i] {
				// 括号不匹配
				return false
			}
		}

	}

	// 如果栈中还有数据,则返回false
	return len(stack) == 0
}

// @lc code=end
