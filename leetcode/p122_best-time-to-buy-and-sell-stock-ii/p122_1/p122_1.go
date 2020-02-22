package p122_1

var MaxProfit = maxProfit

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			return prices[i-1] - min + maxProfit(prices[i:])
		}
	}
	return prices[len(prices)-1] - prices[0]
}
