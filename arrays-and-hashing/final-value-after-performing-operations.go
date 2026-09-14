func finalValueAfterOperations(operations []string) int {
	ans := 0

	for _, op := range operations {
		if op[1] == '+' {
			ans++
		} else {
			ans--
		}
	}

	return ans
}
