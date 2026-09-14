func sortedSquares(nums []int) []int {
    n := len(nums)

    l, r := 0, n - 1

    ans := make([]int, n)

    for i := n - 1; i >= 0; i-- {
        if abs(nums[l]) < abs(nums[r]) {
            ans[i] = nums[r] * nums[r]
            r--
        } else {
            ans[i] = nums[l] * nums[l]
            l++
        }
    }

    return ans
}

func abs (n int) int {
    if n < 0 {
        return -n
    }

    return n
}
