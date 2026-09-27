func searchMatrix(matrix [][]int, target int) bool {
	// this is the exact solution for a normal Binary Search but instead of arr[mid], use matrix[mid/m][mid%m]
	n, m := len(matrix), len(matrix[0]) // num of rows, num of cols
	l, r := 0, m*n-1

	for l <= r {
		mid := l + (r-l)/2
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
