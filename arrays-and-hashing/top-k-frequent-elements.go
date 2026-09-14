func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	bkt := make([][]int, len(nums)+1)

	for key, val := range freq {
		bkt[val] = append(bkt[val], key)
	}

	ans := make([]int, 0, k)

	for i := len(bkt) - 1; i >= 0; i-- {
		if bkt[i] != nil {
			ans = append(ans, bkt[i]...)
			k -= len(bkt[i])
			if k == 0 { break }
		}
	}

	return ans
}
