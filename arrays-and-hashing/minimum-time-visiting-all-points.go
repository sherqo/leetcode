func minTimeToVisitAllPoints(points [][]int) int {
    ans := 0

    for i := 1; i < len(points); i++{
        diffX := abs(points[i][0] - points[i - 1][0])
        diffY := abs(points[i][1] - points[i - 1][1])

        ans += max(diffX, diffY)
    }

    return ans
}
