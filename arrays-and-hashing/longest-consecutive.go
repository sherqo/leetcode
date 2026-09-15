func longestConsecutive(nums []int) int {
	st := make(map[int]struct{})

	for _, num := range nums {
		st[num] = struct{}{}
	}

	ans := 0

	for key := range st {
		if _, exists := st[key-1]; !exists {
			curr := key
			seq := 1

			for {
				if _, exists := st[curr+1]; exists {
					curr++
					seq++
				} else {
					break
				}
			}

			if seq > ans {
				ans = seq
			}
		}
	}

	return ans
}