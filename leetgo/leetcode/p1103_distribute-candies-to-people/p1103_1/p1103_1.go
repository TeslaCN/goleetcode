package p1103_1

func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	pos, candy := 0, 1
	for candies > 0 {
		if candies < candy {
			res[pos] += candies
			break
		}
		res[pos] += candy
		candies -= candy
		pos = (pos + 1) % num_people
		candy++
	}
	return res
}
