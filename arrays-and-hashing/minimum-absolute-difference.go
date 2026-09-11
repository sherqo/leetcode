func minimumAbsDifference(arr []int) [][]int {
	// my initial idea is to sort the array and loop on it to know the abs min diff
	// then loop again to append any two elements with this diff
	// I have a more clever solution with couting sort but nah, it would be too complex
	// and won't achieve a better complexity in practical cases
	// in this solution, I did it in a single pass instead

	slices.Sort(arr)

	ans := make([][]int, 0)
	minDiff := arr[1] - arr[0] // no need for abs since arr[j] > arr[i] if j > i

	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]

		if diff > minDiff {
			continue
		}

		if diff < minDiff {
			minDiff = diff
			ans = make([][]int, 0)
		}

		ans = append(ans, []int{arr[i-1], arr[i]})
	}

	return ans
}

