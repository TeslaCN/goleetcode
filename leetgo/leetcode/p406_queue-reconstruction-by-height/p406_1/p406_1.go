package p406_1

import "sort"

var ReconstructQueue = reconstructQueue

type PeopleInterface [][]int

func (p PeopleInterface) Len() int {
	return len(p)
}

func (p PeopleInterface) Less(i, j int) bool {
	iOrder := p[i][1]
	jOrder := p[j][1]
	iHeight := p[i][0]
	jHeight := p[j][0]
	return iOrder < jOrder || (iOrder == jOrder && iHeight < jHeight)
}

func (p PeopleInterface) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func reconstructQueue(people [][]int) [][]int {
	if len(people) < 2 {
		return people
	}
	sort.Sort(PeopleInterface(people))
	res := make([][]int, 0)
	for i := 0; i < len(people); i++ {
		if people[i][1] == 0 {
			res = append(res, people[i])
			continue
		}
		remain := people[i][1]
		pos := 0
		for remain > 0 || (pos < len(res) && people[i][0] > res[pos][0]) {
			if res[pos][0] >= people[i][0] {
				remain--
			}
			pos++
		}
		res = append(res[:pos], append([][]int{people[i]}, res[pos:]...)...)
	}
	return res
}
