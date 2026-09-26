func evalRPN(tokens []string) int {
	st := make([]int, 0, len(tokens))

	for _, t := range tokens {
		var calc int

		switch t {
		case "+":
			calc = st[len(st)-2] + st[len(st)-1]
			st = st[:len(st)-2]
		case "-":
			calc = st[len(st)-2] - st[len(st)-1]
			st = st[:len(st)-2]
		case "*":
			calc = st[len(st)-2] * st[len(st)-1]
			st = st[:len(st)-2]
		case "/":
			calc = st[len(st)-2] / st[len(st)-1]
			st = st[:len(st)-2]
		default:
			calc, _ = strconv.Atoi(t)
		}

		st = append(st, calc)
	}

	return st[0]
}
