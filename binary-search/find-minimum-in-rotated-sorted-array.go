func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		if nums[r] > nums[mid] { // keep and go left
			r = mid
		} else { // no keep and go right
			l = mid + 1
		}
	}

	return nums[l]
}
