package p146_1

type Node struct {
	key       int
	value     int
	pre, next *Node
}

// Runtime: O(1)
// Space: O(n)
type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	node, exists := this.cache[key]
	if !exists {
		// key 不存在
		return -1
	}
	switch {
	case node.pre != nil && node.next != nil:
		// 节点在中间
		node.pre.next = node.next
		node.next.pre = node.pre
		this.head.pre = node
		node.pre = nil
		node.next = this.head
		this.head = node

	case node.pre != nil:
		// 节点是尾节点
		this.tail = node.pre
		this.tail.next = nil
		this.head.pre = node
		node.pre = nil
		node.next = this.head
		this.head = node

	case node.next != nil:
	default:
		// 节点已在头部或是唯一节点
	}
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}
	if _, exists := this.cache[key]; exists {
		// 节点已存在，更新
		this.cache[key].value = value
		// 更新节点位置
		_ = this.Get(key)
		return
	}
	if len(this.cache) >= this.capacity {
		// 缓存已满
		tailNode := this.tail
		if len(this.cache) == 1 {
			// 唯一节点
			this.head = nil
			this.tail = nil
		} else {
			// 移除最后一个节点
			this.tail = tailNode.pre
			this.tail.next = nil
			tailNode.pre = nil
		}
		// 删除节点
		delete(this.cache, tailNode.key)
	}
	newNode := &Node{
		key:   key,
		value: value,
		next:  this.head,
	}
	this.cache[key] = newNode
	if this.head == nil {
		// empty
		this.tail = newNode
	} else {
		this.head.pre = newNode
	}
	this.head = newNode
}

/*
执行用时 : 132 ms , 在所有 Go 提交中击败了 91.18% 的用户
内存消耗 : 15 MB , 在所有 Go 提交中击败了 88.24% 的用户
*/
