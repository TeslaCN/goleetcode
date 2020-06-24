package p406_2

import "sort"

var ReconstructQueue = reconstructQueue

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		iHeightDesc := -people[i][0]
		jHeightDesc := -people[j][0]
		return iHeightDesc < jHeightDesc || (iHeightDesc == jHeightDesc && people[i][1] < people[j][1])
	})
	res := make([][]int, 0)
	for _, person := range people {
		res = append(res[:person[1]], append([][]int{person}, res[person[1]:]...)...)
	}
	return res
}
