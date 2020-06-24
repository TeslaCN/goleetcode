package p460_lfu_cache

import (
	"github.com/TeslaCN/goleetcode/leetcode/p460_lfu-cache/p460_1"
	"github.com/TeslaCN/goleetcode/leetcode/p460_lfu-cache/p460_2"
	"testing"
)

func Test_1_1(t *testing.T) {
	cache := p460_1.Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(1))
	t.Log(cache.Get(2))
	cache.Put(3, 3)
	t.Log(cache.Get(2))
	t.Log(cache.Get(3))
	cache.Put(4, 4)
	t.Log(cache.Get(1))
	t.Log(cache.Get(3))
	t.Log(cache.Get(4))
}

func Test_2_1(t *testing.T) {
	cache := p460_2.Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(1) == 1)
	t.Log(cache.Get(2) == 2)
	cache.Put(3, 3)
	t.Log(cache.Get(1) == -1)
	t.Log(cache.Get(2) == 2)
	t.Log(cache.Get(3) == 3)
	cache.Put(4, 4)
	t.Log(cache.Get(1) == -1)
	t.Log(cache.Get(3) == -1)
	t.Log(cache.Get(4) == 4)
}
func Test_2_2(t *testing.T) {
	cache := p460_2.Constructor(0)
	cache.Put(0, 0)
	t.Log(cache.Get(0) == -1)
}

func Test_2_3(t *testing.T) {
	/*
		["LFUCache","put","put","get","put","put","get"]
		[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
	*/
	cache := p460_2.Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(2) == 2)
	cache.Put(1, 1)
	cache.Put(4, 1)
	t.Log(cache.Get(2) == 2)
}

func Test_2_4(t *testing.T) {
	/*
	   ["LFUCache","put","put","put","put","get"]
	   [[2],[3,1],[2,1],[2,2],[4,4],[2]]
	*/
	cache := p460_2.Constructor(2)
	cache.Put(3, 1)
	cache.Put(2, 1)
	cache.Put(2, 2)
	cache.Put(4, 4)
	t.Log(cache.Get(2) == 2)
}

func Test_2_5(t *testing.T) {
	/*
	   ["LFUCache","put","put","get","put","put","get"]
	   [[2],[2,1],[2,2],[2],[1,1],[2,1],[2]]
	*/
	cache := p460_2.Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(2) == 2)
	cache.Put(1, 1)
	cache.Put(2, 1)
	t.Log(cache.Get(2) == 1)
}

func Test_2_6(t *testing.T) {
	/*
		["LFUCache","put","put","put","put","get","get"]
		[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	*/
	cache := p460_2.Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 2)
	cache.Put(2, 3)
	cache.Put(4, 1)
	t.Log(cache.Get(1) == -1)
	t.Log(cache.Get(2) == 3)
}

/*
["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
*/
