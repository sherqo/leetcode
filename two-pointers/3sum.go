func threeSum(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] { // for dups
			continue
		}

		l, r := i+1, n-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 { // make it bigger
				l++
			} else if sum > 0 { // make it smaller
				r--
			} else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})

				for l < r && nums[l] == nums[l+1] { // skip
					l++
				}

				for l < r && nums[r-1] == nums[r] { // skip
					r--
				}

				l++
				r--
			}
		}
	}

	return ans
}
