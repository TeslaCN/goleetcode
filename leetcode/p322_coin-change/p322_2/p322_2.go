package p322_2

var CoinChange = coinChange

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	res := make([]int, amount)
	for i := 1; i <= amount; i++ {
		res[i-1] = c(coins, i, res)
	}
	return res[amount-1]
}

func c(coins []int, amount int, res []int) int {
	if amount == 0 {
		return 0
	}
	minValue := -1
	for _, coin := range coins {
		if amount == coin {
			// min[amount] = 1
			minValue = 0
		} else if amount < coin {
			continue
		} else {
			value := res[amount-coin-1]
			if value > 0 && (minValue < 0 || value < minValue) {
				minValue = value
			}
		}
	}
	if minValue >= 0 {
		return minValue + 1
	} else {
		return minValue
	}
}
