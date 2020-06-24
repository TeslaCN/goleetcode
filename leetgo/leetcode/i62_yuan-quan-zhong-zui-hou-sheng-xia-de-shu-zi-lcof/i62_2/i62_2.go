package i62_2

// Deprecated
var LastRemaining = lastRemaining

type node struct {
	val  int
	next *node
}

// Deprecated: slow
func lastRemaining(n int, m int) int {
	root := &node{val: 0}
	p := root
	for i := 1; i < n; i++ {
		p.next = &node{val: i}
		p = p.next
	}
	p.next = root

	for n > 1 {
		for i := 0; i < m-1; i++ {
			p = p.next
		}
		p.next = p.next.next
		n--
	}
	return p.val
}
