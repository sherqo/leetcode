func dailyTemperatures(temperatures []int) []int {
	st := make([]int, 0, len(temperatures)) // preallocating the stack is good :) 
	ans := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for len(st) > 0 && temp > temperatures[st[len(st)-1]] {
			ans[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}

		st = append(st, i)
	}

	return ans
}
