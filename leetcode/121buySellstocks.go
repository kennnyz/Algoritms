package leetcode

func buyAndSell(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func maxProfit2(prices []int) int {
	buyPrices := make([]int, len(prices))

	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			buyPrices[i] = -1
		}
	}
	for j := 0; j < len(prices); j++ {
		if buyPrices[j] != 0 {
			profit += prices[j+1] - prices[j]

		}
	}
	return profit
}
