func maxProfit(prices []int) int {
	// we need to get a big number from the right with a small number from the left, that's it

	prefix := make([]int, len(prices)) // to keep the prefix minimums

	prefix[0] = prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < prefix[i-1] {
			prefix[i] = prices[i]
		} else {
			prefix[i] = prefix[i-1]
		}
	}

	ans := 0

	for i := len(prices) - 1; i > 0; i-- {
		cur := prices[i] - prefix[i-1]

		if cur > ans {
			ans = cur
		}
	}

	return ans
}

func maxProfit(prices []int) int {
	// we need to get a big number from the right with a small number from the left, that's it

	buy, ans := prices[0], 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else if prices[i]-buy > ans {
			ans = prices[i] - buy
		}
	}

	return ans
}
