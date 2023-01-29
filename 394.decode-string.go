/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 *
 * https://leetcode.cn/problems/decode-string/description/
 *
 * algorithms
 * Medium (56.86%)
 * Likes:    1394
 * Dislikes: 0
 * Total Accepted:    219.7K
 * Total Submissions: 386.4K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * 给定一个经过编码的字符串，返回它解码后的字符串。
 *
 * 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
 *
 * 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
 *
 * 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "3[a]2[bc]"
 * 输出："aaabcbc"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "3[a2[c]]"
 * 输出："accaccacc"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "2[abc]3[cd]ef"
 * 输出："abcabccdcdcdef"
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "abc3[cd]xyz"
 * 输出："abccdcdcdxyz"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 30
 * s 由小写英文字母、数字和方括号 '[]' 组成
 * s 保证是一个 有效 的输入。
 * s 中所有整数的取值范围为 [1, 300]
 *
 *
 */
package main

import (
	"strconv"
	"strings"
	"unicode"
)

// @lc code=start
func decodeString(s string) string {
	// 将字符串加入栈中 遇到']'之后出栈
	// 输入：s = "3[a]2[bc]"
	stack := make([]string, 0)
	ptr := 0
	// s := "3[a]2[bc]"
	for ptr < len(s) {
		cur := s[ptr]
		// 三种情况 数字 [ 字母 ]
		if unicode.IsDigit(rune(cur)) {
			// 数字 可能为两位或三位 这里需要特殊处理
			digits := GetDigits(s, &ptr)
			// 直接入栈
			stack = append(stack, digits)
		} else if unicode.IsLetter(rune(cur)) || cur == '[' {
			// 字母 入栈
			stack = append(stack, string(cur))
			ptr++
		} else {
			ptr++
			// ] 出栈
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sb := strings.Builder{}

			// 拿到所有字母
			for pop != "[" {
				sb.WriteString(pop)
				pop = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}

			// 将字母反转 因为从栈中获取的数据是反方向的
			revers := ReverseString(sb.String())

			// 拿数字
			digit, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]

			// 将数字和字母结合 放入栈中
			sub := StringWithDigits(revers, digit)
			stack = append(stack, sub)
		}
	}

	return GetStringWithStacks(stack)
}

func GetDigits(s string, ptr *int) string {
	digit := s[*ptr]
	sb := strings.Builder{}
	// 只要是数字 就一直遍历 并加入到sb中
	for unicode.IsDigit(rune(digit)) {
		sb.WriteByte(digit)
		*ptr++
		digit = s[*ptr]
	}

	return sb.String()
}

// 反转字符串 通过string builder
func ReverseStringWithBuilder(s string) string {
	// 转化为rune可以存放一些中文
	runes := []rune(s)
	sb := strings.Builder{}
	// 获取每一位字符粗
	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])
	}

	return sb.String()
}

// 反转字符串 不使用额外空间
func ReverseString(s string) string {
	runes := []rune(s)
	for start, end := 0, len(runes)-1; start < end; start, end = start+1, end-1 {
		// 交换位置
		runes[start], runes[end] = runes[end], runes[start]
	}

	return string(runes)
}

func StringWithDigits(sub string, digit int) string {
	sb := strings.Builder{}
	for i := 0; i < digit; i++ {
		sb.WriteString(sub)
	}
	return sb.String()
}

func GetStringWithStacks(stack []string) string {
	sb := strings.Builder{}
	for _, v := range stack {
		sb.WriteString(v)
	}

	return sb.String()
}

// @lc code=end
// func main() {
// 	// Testcase: "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
// 	// Answer:          "zzzyypqfejkjkfejkjkfejkjkfejkjkyypqfejkjkfejkjkfejkjkfejkjkef"
// 	// Expected Answer: "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"
// 	s := "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
// 	fmt.Printf("decodeString(s): %v\n", decodeString(s))
// }
