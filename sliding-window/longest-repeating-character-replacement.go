func characterReplacement(s string, k int) int {
	ans := 0
	for i := byte('A'); i <= byte('Z'); i++ {
		if pot := getAnsForSingleLetter(i, s, k); pot > ans {
			ans = pot
		}
	}

	return ans
}

func getAnsForSingleLetter(letter byte, s string, k int) int {
	// let's make it for just a single letter B and then loop on them all
	l := 0
	ans := 0
	freq := 0 // count of non-'X' characters; problem if freq > k
	for r := 0; r < len(s); r++ {
		// do the thing
		if s[r] != letter {
			freq++
		}

		// fix until fixed
		for freq > k {
			if s[l] != letter {
				freq--
			}
			l++
		}

		// calc stuff
		if win := r - l + 1; win > ans {
			ans = win
		}
	}

	return ans
}
