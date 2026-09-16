func minSubArrayLen(target int, nums []int) int {
	// so we're optimizing for 2 things, make the sum great above the target
	// and also make the length small as possible
	// I will make the right to grow and maintain the left on demand

	ans := len(nums) + 1
	sum := 0 // current sum, not the max
	l := 0

	for r := 0; r < len(nums); r++ {
		sum += nums[r] // add the new element

		for sum >= target { // as long as we're great, shrink from the left
			if (r - l + 1) < ans {
				ans = r - l + 1
			}

			sum -= nums[l]
			l++
		}
	}

	if ans > len(nums) {
		return 0
	}
	return ans
}
