package p322_1

var CoinChange = coinChange

var min map[int]int

func coinChange(coins []int, amount int) int {
	min = make(map[int]int)
	return c(coins, amount)
}

func c(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if value, exists := min[amount]; exists {
		return value
	}
	minValue := -1
	for _, coin := range coins {
		if amount == coin {
			// min[amount] = 1
			minValue = 0
		} else if amount < coin {
			continue
		} else {
			value := coinChange(coins, amount-coin)
			if value > 0 && (minValue < 0 || value < minValue) {
				minValue = value
			}
		}
	}
	if minValue >= 0 {
		min[amount] = minValue + 1
		return minValue + 1
	} else {
		min[amount] = minValue
		return minValue
	}
}
