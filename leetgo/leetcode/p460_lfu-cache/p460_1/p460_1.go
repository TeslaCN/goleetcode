package p460_1

type LFUCache struct {
	cap     int
	cache   map[int]int
	times   map[int]int
	freq    map[int]map[int]struct{}
	minFreq int
}

// Deprecated: 能够实现 LFU，但相同频率元素无法保证 LRU
func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		cache:   make(map[int]int),
		times:   make(map[int]int),
		freq:    make(map[int]map[int]struct{}),
		minFreq: 0,
	}
}

func (this *LFUCache) Get(key int) int {
	value, exists := this.cache[key]
	if !exists {
		return -1
	}
	this.times[key]++
	delete(this.freq[this.times[key]-1], value)

	_, exists = this.freq[this.times[key]]
	if !exists {
		this.freq[this.times[key]] = make(map[int]struct{})
	}
	this.freq[this.times[key]][value] = struct{}{}
	if this.times[key]-1 == this.minFreq && len(this.freq[this.times[key]-1]) == 0 {
		this.minFreq = this.times[key]
	}
	return value
}

func (this *LFUCache) Put(key int, value int) {
	if len(this.cache) == this.cap {
		var oldKey int
		for k := range this.freq[this.minFreq] {
			oldKey = k
			break
		}
		delete(this.freq[this.minFreq], oldKey)
		delete(this.cache, oldKey)
		delete(this.times, oldKey)
	}
	this.cache[key] = value
	this.minFreq = 0
	this.times[key] = 0
	if _, exists := this.freq[0]; !exists {
		this.freq[0] = make(map[int]struct{})
	}
	this.freq[0][key] = struct{}{}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
