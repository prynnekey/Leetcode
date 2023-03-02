/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 *
 * https://leetcode.cn/problems/implement-queue-using-stacks/description/
 *
 * algorithms
 * Easy (68.67%)
 * Likes:    828
 * Dislikes: 0
 * Total Accepted:    313.6K
 * Total Submissions: 456.7K
 * Testcase Example:  '["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
 *
 * 实现 MyQueue 类：
 *
 *
 * void push(int x) 将元素 x 推到队列的末尾
 * int pop() 从队列的开头移除并返回元素
 * int peek() 返回队列开头的元素
 * boolean empty() 如果队列为空，返回 true ；否则，返回 false
 *
 *
 * 说明：
 *
 *
 * 你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty
 * 操作是合法的。
 * 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：
 * ["MyQueue", "push", "push", "peek", "pop", "empty"]
 * [[], [1], [2], [], [], []]
 * 输出：
 * [null, null, null, 1, 1, false]
 *
 * 解释：
 * MyQueue myQueue = new MyQueue();
 * myQueue.push(1); // queue is: [1]
 * myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
 * myQueue.peek(); // return 1
 * myQueue.pop(); // return 1, queue is [2]
 * myQueue.empty(); // return false
 *
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= x <= 9
 * 最多调用 100 次 push、pop、peek 和 empty
 * 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。
 *
 *
 */
package main

import "sync"

// @lc code=start
type MyQueue struct {
	Stack []int
	Back  []int
	Lock  sync.RWMutex
}

// 和146题冲突
// func Constructor() MyQueue {
// 	return MyQueue{
// 		Stack: make([]int, 0),
// 		Back:  make([]int, 0),
// 	}
// }

func (m *MyQueue) Push(x int) {
	m.Lock.Lock()
	// 将back中的所有元素先放入stack中
	for len(m.Back) != 0 {
		// 取出最后一位
		val := m.Back[len(m.Back)-1]
		// 将最后一位删除
		m.Back = m.Back[:len(m.Back)-1]

		// 将取出的元素放入stack中
		m.Stack = append(m.Stack, val)
	}
	m.Stack = append(m.Stack, x)
	m.Lock.Unlock()
}

func (m *MyQueue) Pop() int {
	m.Lock.Lock()
	// 将栈中的元素全部移动到Back中
	for len(m.Stack) != 0 {
		// 获取最后一位元素
		val := m.Stack[len(m.Stack)-1]
		// 将最后一位元素删除
		m.Stack = m.Stack[:len(m.Stack)-1]

		// 将val放入Back中
		m.Back = append(m.Back, val)
	}
	m.Lock.Unlock()

	// 如果back中没有元素 说明当前栈为空 返回无数据
	if len(m.Back) == 0 {
		return 0
	}

	// 有数据 将最后一位数据取出来
	val := m.Back[len(m.Back)-1]

	m.Lock.Lock()
	// 将最后一位数据删除
	m.Back = m.Back[:len(m.Back)-1]
	m.Lock.Unlock()

	return val
}

func (m *MyQueue) Peek() int {
	m.Lock.Lock()
	// 将stack中的元素放入back中
	for len(m.Stack) != 0 {
		val := m.Stack[len(m.Stack)-1]
		m.Stack = m.Stack[:len(m.Stack)-1]
		m.Back = append(m.Back, val)
	}
	m.Lock.Unlock()
	if len(m.Back) == 0 {
		return 0
	}

	// 取出back的最后一位
	return m.Back[len(m.Back)-1]
}

func (m *MyQueue) Empty() bool {
	// 判断back和stack同时为空才为空
	return len(m.Back) == 0 && len(m.Stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
