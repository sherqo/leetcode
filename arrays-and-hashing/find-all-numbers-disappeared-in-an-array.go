
func findDisappearedNumbers(nums []int) []int {
	// the logic will be like:
	// loop on the whole array, and mark the integer you
	// see to seen, but instead of making another data
	// structure for this, use the same array nums
	// by saving this info in a negative index
	// after doing all of this, loop on the whole array
	// again and save the positive indexes only

	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		idx := num - 1

		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	ans := []int{}
	for idx, num := range nums {
		if num > 0 {
			ans = append(ans, idx+1)
		}
	}

	return ans
}
