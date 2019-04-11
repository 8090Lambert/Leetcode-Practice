package maxProfit

func MaxProfit(prices []int) int {
	count := len(prices)
	if count <= 0 {
		return 0
	}
	earning := 0
	min := prices[0]
	for index := 1; index < count; index++ {
		if prices[index] < min {
			min = prices[index]
		} else if prices[index]-min > earning {
			earning = prices[index] - min
		}
	}

	return earning
}

func MaxProfit1(prices []int) int {
	count := len(prices)
	if count <= 0 {
		return 0
	}

	earning := 0
	max := 0
	min := 0
	for index := 1; index < count; index++ {
		if prices[index] < prices[min] {
			min = index
		}
		if prices[index] > prices[min] {
			max = index
		}

		if max > min {
			earning += prices[max] - prices[min]
			max = index
			min = index
		}
	}

	//fmt.Println(earning)
	return earning
}
