package _22_maxProfit

func maxProfit(prices []int) int {
	earn, min, max := 0, 0, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[min] {
			min = i
		} else {
			max = i
		}
		if max > min && prices[max] > prices[min] {
			earn += prices[max] - prices[min]
			min, max = i, i
		}
	}
	return earn
}

func maxProfitOther(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i]-prices[i-1]
		}
	}
	return res
}
