/*
* @lc app=leetcode.cn id=146 lang=golang
*
* [146] LRU 缓存
*
* https://leetcode.cn/problems/lru-cache/description/
*
  - algorithms
  - Medium (53.43%)
  - Likes:    2566
  - Dislikes: 0
  - Total Accepted:    452.6K
  - Total Submissions: 846.8K
  - Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
    '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'

*
* 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
*
* 实现 LRUCache 类：
*
*
*
*
* LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
* int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
* void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
* key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
*
*
* 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*
*
*
*
*
* 示例：
*
*
* 输入
* ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
* [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
* 输出
* [null, null, null, 1, null, -1, null, -1, 3, 4]
*
* 解释
* LRUCache lRUCache = new LRUCache(2);
* lRUCache.put(1, 1); // 缓存是 {1=1}
* lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
* lRUCache.get(1);    // 返回 1
* lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
* lRUCache.get(2);    // 返回 -1 (未找到)
* lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
* lRUCache.get(1);    // 返回 -1 (未找到)
* lRUCache.get(3);    // 返回 3
* lRUCache.get(4);    // 返回 4
*
*
*
*
* 提示：
*
*
* 1 <= capacity <= 3000
* 0 <= key <= 10000
* 0 <= value <= 10^5
* 最多调用 2 * 10^5 次 get 和 put
*
*
*/
package main

// @lc code=start

// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
// get 获取节点 并删除 然后插入到头部 -> 链表 hashmap
// put 插入到头部，如果存在 删除，如果不存在 删除尾部元素 -> 双向链表

type Node struct {
	Key   int
	Value int
	Next  *Node
	Pre   *Node
}

type LRUCache struct {
	Capacity int
	Head     *Node
	Tail     *Node
	Hashmap  map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	// head <-> tail
	head.Next = tail
	tail.Pre = head

	return LRUCache{
		Capacity: capacity,
		Head:     head,
		Tail:     tail,
		Hashmap:  make(map[int]*Node),
	}
}

func (lru *LRUCache) Get(key int) int {
	// 判断是否存在
	if node, ok := lru.Hashmap[key]; ok {
		remove(node)
		insert(lru.Head, node)
		return node.Value
	}

	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.Hashmap[key]; ok {
		// 存在
		node.Value = value
		remove(node)
		insert(lru.Head, node)
	} else {
		node := &Node{
			Key:   key,
			Value: value,
		}
		lru.Hashmap[key] = node
		insert(lru.Head, node)
		// 判断是否满了
		if len(lru.Hashmap) > lru.Capacity {
			// 删除尾部元素
			tail := lru.Tail.Pre
			remove(tail)
			delete(lru.Hashmap, tail.Key)
		}
	}
}

func remove(node *Node) {
	// pre <-> node <-> next
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func insert(head, node *Node) {
	next := head.Next
	// head <-> node <-> ...node <-> tail

	// head -> node -> next
	head.Next = node
	node.Next = next

	// head <-> node <-> next
	node.Pre = head
	next.Pre = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
