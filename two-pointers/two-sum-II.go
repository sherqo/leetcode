func twoSum(nums []int, target int) []int {
	p1, p2 := 0, len(nums)-1

	for p1 < p2 {
		sum := nums[p1] + nums[p2]
		if sum == target {
			return []int{p1 + 1, p2 + 1}
		} else if sum > target {
			p2--
		} else {
			p1++
		}
	}

	return []int{p1 + 1, p2 + 1}
}