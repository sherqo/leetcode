func lengthOfLongestSubstring(s string) int {
	var freq [256]int

	ans := 0
	l := 0

	for r := 0; r < len(s); r++ {
		freq[s[r]]++

		for freq[s[r]] > 1 {
			freq[s[l]]--
			l++
		}

		if win := r - l + 1; win > ans {
			ans = win
		}
	}

	return ans
}
