package p122_2

var MaxProfit = maxProfit

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}
