package maxProfit

func maxProfit(prices []int) int {
	count := len(prices)
	min := prices[0]
	earning := 0
	for i := 1; i < count; i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i] - min > earning {
			earning = prices[i] - min
		}
	}
	return earning
}

func maxProfit2(prices []int) int {
	count := len(prices)
	min, max, earning := 0, 0, 0
	for i := 1; i < count; i++ {
		if prices[i] < prices[min] {
			min = i
		} else {
			max = i
		}
		if max > min && prices[max] - prices[min] > 0 {
			earning += prices[max]-prices[min]
			max, min = i, i
		}
	}
	return earning
}