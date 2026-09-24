func timeRequiredToBuy(tickets []int, k int) int {
    // I got this clever solution by doing some simulations on a whiteboard

    ans := 0
    for i, t := range tickets {
        if i <= k {
            ans += min(t, tickets[k])
        } else {
            ans += min(t, tickets[k]-1)
        }
    }
    return ans
}
