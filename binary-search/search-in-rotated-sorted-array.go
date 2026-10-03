func search(a []int, t int) int {
	l, r := 0, len(a)-1

	for l <= r {
		mid := l + (r-l)/2

		if a[mid] == t {
			return mid
		}

		if a[l] <= a[mid] { // the left part is sorted
			if t < a[mid] && t >= a[l] { // go left
				r = mid - 1
			} else { // go right
				l = mid + 1
			}
		} else {
			if t > a[mid] && t <= a[r] { // go right
				l = mid + 1
			} else { // go left
				r = mid - 1
			}
		}
	}

	return -1
}
