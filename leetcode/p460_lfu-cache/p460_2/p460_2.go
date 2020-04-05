package p460_2

type LFUCache struct {
	cap     int
	minFreq int
	cache   map[int]*Node
	freq    map[int][2]*Node
}

type Node struct {
	key       int
	value     int
	freq      int
	pre, next *Node
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		minFreq: 1,
		cache:   make(map[int]*Node),
		freq:    make(map[int][2]*Node),
	}
}

func (this *LFUCache) Get(key int) int {
	node, exists := this.cache[key]
	if !exists {
		return -1
	}
	if node.pre == nil && node.next == nil {
		// 该频率唯一节点
		delete(this.freq, node.freq)
		if this.minFreq == node.freq {
			this.minFreq++
		}
	} else {
		if node.pre != nil && node.next != nil {
			// 该频率有多个节点，删掉双向链表中的指定节点
			node.pre.next = node.next
			node.next.pre = node.pre
		} else if node.pre != nil {
			// 该结点为双向链表尾节点
			node.pre.next = nil
			if node.next == nil {
				this.freq[node.freq] = [2]*Node{this.freq[node.freq][0], node.pre}
			}
		} else if node.next != nil {
			// 该结点为双向链表头节点
			node.next.pre = nil
			if node.pre == nil {
				this.freq[node.freq] = [2]*Node{node.next, this.freq[node.freq][1]}
			}
		}
		node.pre, node.next = nil, nil
	}
	if headTail, exists := this.freq[node.freq+1]; !exists {
		// 新的频率还没有节点
		this.freq[node.freq+1] = [2]*Node{node, node}
	} else {
		// 插入新频率的链表头
		headTail[0].pre = node
		node.next = headTail[0]
		headTail[0] = node
		this.freq[node.freq+1] = headTail
	}
	node.freq++
	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	// key 已存在，更新 value 并增加一次访问
	if _, exists := this.cache[key]; exists {
		this.cache[key].value = value
		this.Get(key)
		return
	}

	// 缓存满，需要清理
	if len(this.cache) == this.cap {
		headTail := this.freq[this.minFreq]
		node := headTail[1]

		if headTail[0] == headTail[1] {
			// 该频率只有一个数据，直接删掉该频率的key
			headTail[0], headTail[1] = nil, nil
			delete(this.freq, this.minFreq)
		} else {
			// 删掉末尾节点
			headTail[1] = node.pre
			node.pre = nil
			headTail[1].next = nil
			this.freq[this.minFreq] = headTail
		}
		delete(this.cache, node.key)
	}
	this.minFreq = 1
	this.cache[key] = &Node{
		key:   key,
		value: value,
		freq:  this.minFreq,
	}
	headTail, exists := this.freq[this.minFreq]
	if !exists {
		// 该频率没有节点
		this.freq[this.minFreq] = [2]*Node{this.cache[key], this.cache[key]}
	} else {
		// 插入新频率的链表头
		headTail[0].pre = this.cache[key]
		this.cache[key].next = headTail[0]
		headTail[0] = this.cache[key]
		this.freq[this.minFreq] = headTail
	}
}

/*
执行用时 : 148 ms , 在所有 Go 提交中击败了 59.70% 的用户
内存消耗 : 17.2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
