package p146_1

import "testing"

func TestLRUCache_1_1(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(1) == 1)
	cache.Put(3, 3)
	t.Log(cache.Get(2) == -1)
	cache.Put(4, 4)
	t.Log(cache.Get(1) == -1)
	t.Log(cache.Get(3) == 3)
	t.Log(cache.Get(4) == 4)
}

func TestLRUCache_1_2(t *testing.T) {
	cache := Constructor(0)
	cache.Put(1, 1)
	t.Log(cache.Get(1) == -1)
}

func TestLRUCache_1_3(t *testing.T) {
	cache := Constructor(1)
	cache.Put(1, 1)
	t.Log(cache.Get(1) == 1)
	cache.Put(2, 2)
	t.Log(cache.Get(1) == -1)
	t.Log(cache.Get(2) == 2)
	cache.Put(2, 22)
	t.Log(cache.Get(2) == 22)
}

/*
["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
*/
func TestLRUCache_1_4(t *testing.T) {
	cache := Constructor(10)
	cache.Put(10, 13)
	cache.Put(3, 17)
	cache.Put(6, 11)
	cache.Put(10, 5)
	cache.Put(9, 10)
	t.Log(cache.Get(13) == -1)
	cache.Put(2, 19)
	t.Log(cache.Get(2) == 19)
	t.Log(cache.Get(3) == 17)

}

func TestLRUCache_1_5(t *testing.T) {
	cache := Constructor(3)
	for i := 1; i <= 3; i++ {
		cache.Put(i, i)
	}
	for i := 1; i <= 3; i++ {
		cache.Get(i)
	}
	cache.Get(1)
}

func Test_1(t *testing.T) {
	cases := []struct {
		method []string
		args   [][]int
	}{
		{[]string{
			"LRUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put",
		}, [][]int{
			{10}, {10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26},
		}},
	}

	for _, c := range cases {
		cache := Constructor(c.args[0][0])
		for i := 1; i < len(c.method); i++ {
			switch c.method[i] {
			case "put":
				t.Logf("Put %v -> %v", c.args[i][0], c.args[i][1])
				cache.Put(c.args[i][0], c.args[i][1])
			case "get":
				t.Logf("Get %v -> %v", c.args[i][0], cache.Get(c.args[i][0]))
			}
		}
	}
}
