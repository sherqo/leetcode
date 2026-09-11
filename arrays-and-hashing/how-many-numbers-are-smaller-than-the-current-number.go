func smallerNumbersThanCurrent(nums []int) []int {
    count := make([]int, 101)
    
    for _, num := range nums {
        count[num]++
    }
    
    prev := count[0]
    count[0] = 0
    for i := 1; i < 101; i++ {
        count[i], prev = count[i-1]+prev, count[i]
    }
    
    ans := make([]int, len(nums))
    for i, num := range nums {
        ans[i] = count[num]
    }
    
    return ans
}

