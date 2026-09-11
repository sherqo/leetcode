func containsNearbyDuplicate(nums []int, k int) bool {
	// my initial idea is to make a list or a set with size of k
	// and for each element in the array: check if it exists in this
	// short list or not; but idk how to do this in golang so I will
	// make a map with the size of the array we have

	mp := make(map[int]int, len(nums)) // takes integer and returns and index

	for idx, num := range nums {
		if idx2, exists := mp[num]; exists {
			if idx-idx2 <= k {
				return true
			}
		}

		mp[num] = idx
	}

	return false
}


