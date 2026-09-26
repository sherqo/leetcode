func generateParenthesis(n int) []string {
	var ans []string

	buffer := make([]byte, 0, 2*n)

	var bt func(open int, close int)
	bt = func(open int, close int) {
		if len(buffer) == 2*n { // string length reaches 2*n
			ans = append(ans, string(buffer))
			return
		}

		if open < n {
			buffer = append(buffer, '(')
			bt(open+1, close)
			buffer = buffer[:len(buffer)-1]
		}

		if close < open {
			buffer = append(buffer, ')')
			bt(open, close+1)
			buffer = buffer[:len(buffer)-1]
		}
	}

	bt(0, 0) // backtrack
	return ans
}
