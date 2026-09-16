func maxArea(height []int) int {
	// the rule: the area between any to indexes i and j:
	// (j - i) * min(h[i], h[j]) => need to find the i, j to get the max

	l, r, ans := 0, len(height)-1, 0

	for l < r {
		cur := (r - l) * height[l]

		if height[l] < height[r] {
			l++
		} else {
			cur = (r - l) * height[r]
			r--
		}

		if cur > ans {
			ans = cur
		}
	}

	return ans
}
