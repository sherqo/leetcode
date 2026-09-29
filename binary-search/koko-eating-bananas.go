func minEatingSpeed(piles []int, h int) int {
	isEnough := func(k int) bool {
		hours := 0

		for _, amount := range piles {
			hours += (amount + k - 1) / k // same as: roundup(amount/k)

			if hours > h {
				return false
			}
		}

		return hours <= h
	}

	l, r := 1, 0
	for _, amount := range piles {
		if amount > r {
			r = amount
		}
	}

	for l < r {
		k := l + (r-l)/2

		if isEnough(k) { // ok, reduce and keep
			r = k
		} else { // not ok, increase and don't keep
			l = k + 1
		}
	}

	return l
}
