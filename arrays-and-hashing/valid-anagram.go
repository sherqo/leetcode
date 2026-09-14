func isAnagram(s string, t string) bool {
	r1 := []rune(s)
	r2 := []rune(t)

	slices.Sort(r1)
	slices.Sort(r2)

	return string(r1) == string(r2)
}
