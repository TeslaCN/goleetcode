package p1431_1

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max, ans := 0, make([]bool, len(candies))
	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}
	for i := 0; i < len(candies); i++ {
		ans[i] = candies[i]+extraCandies >= max
	}
	return ans
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2.3 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
