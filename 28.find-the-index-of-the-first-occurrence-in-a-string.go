/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 *
 * https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
 *
 * algorithms
 * Medium (42.09%)
 * Likes:    1790
 * Dislikes: 0
 * Total Accepted:    814.7K
 * Total Submissions: 1.9M
 * Testcase Example:  '"sadbutsad"\n"sad"'
 *
 * 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0
 * 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：haystack = "sadbutsad", needle = "sad"
 * 输出：0
 * 解释："sad" 在下标 0 和 6 处匹配。
 * 第一个匹配项的下标是 0 ，所以返回 0 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：haystack = "leetcode", needle = "leeto"
 * 输出：-1
 * 解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= haystack.length, needle.length <= 10^4
 * haystack 和 needle 仅由小写英文字符组成
 *
 *
 */
package main

// @lc code=start

// 暴力BF
// func strStr(haystack string, needle string) int {
// 	m := len(haystack)
// 	n := len(needle)
//
// 	flag := false
//
// 	for i := 0; i <= m-n; i++ {
// 		j := 0
// 		for j = 0; j < n; j++ {
// 			if haystack[i+j] != needle[j] {
// 				flag = false
// 				break
// 			}
// 			flag = true
// 		}
//
// 		// 全部相等
// 		if flag {
// 			return i
// 		}
// 	}
//
// 	return -1
// }

// RK算法:hash算法
// func strStr(haystack string, needle string) int {
// 	m := len(haystack)
// 	n := len(needle)
//
// 	// haystackHash := hash(haystack, 0, n)
// 	var haystackHash byte
// 	needleHash := hash(needle, 0, n)
//
// 	for start := 0; start <= m-n; start++ {
// 		haystackHash = hash(haystack, start, n+start, haystackHash)
// 		if haystackHash == needleHash {
// 			// hash相等，进行字符串比较
//
// 			flag := false
//
// 			for i := 0; i <= m-n; i++ {
// 				j := 0
// 				for j = 0; j < n; j++ {
// 					if haystack[i+j] != needle[j] {
// 						flag = false
// 						break
// 					}
// 					flag = true
// 				}
//
// 				// 全部相等
// 				if flag {
// 					return i
// 				}
// 			}
// 		}
// 	}
//
// 	return -1
// }
//
// // 求字符串的hash
// //
// // previousHash 为可选参数，只能为一个
// func hash(str string, start, end int, previousHash ...byte) (hash byte) {
// 	if len(previousHash) == 0 || start == 0 {
// 		for i := start; i < end; i++ {
// 			hash += str[i]
// 		}
// 	} else {
// 		hash = previousHash[0] - str[start-1] + str[end-1]
// 	}
//
// 	return
// }

// KMP算法
func strStr(haystack string, needle string) int {
	// 获取next数组
	next := getNext(needle)

	i, j := 0, 0

	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else if j > 0 {
			j = next[j-1]
		} else {
			// 说明j==0
			i++
		}

		if j == len(needle) {
			return i - j
		}
	}

	return -1
}

// 获取next数组
func getNext(pattern string) []int {
	length := len(pattern)
	next := make([]int, length)

	i, j := 1, 0

	for i < len(pattern) {
		if pattern[i] == pattern[j] {
			next[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = next[j-1]
		} else {
			// 说明j==0
			i++
		}
	}

	return next
}

// @lc code=end
