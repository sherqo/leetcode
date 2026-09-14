func plusOne(digs []int) []int {
	carry := 0
	n := len(digs)

	digs[n-1] += 1

	if digs[n-1] > 9 {
		carry = 1
		digs[n-1] = 0

		for i := n - 2; i >= 0; i-- {
			res := digs[i] + carry
			if res > 9 {
				carry = 1
                digs[i] = 0
			} else {
				carry = 0
                digs[i] = res
			}
		}

        if carry == 1 {
            digs = append([]int{1}, digs...)
        }
	}

	return digs
}
