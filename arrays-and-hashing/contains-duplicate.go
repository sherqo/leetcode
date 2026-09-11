func containsDuplicate(nums []int) bool {
    seen := make(map[int]struct{}, len(nums)) //add len to pre-allocate

    for _, val := range nums {
        if _, exists := seen[val]; exists { //single line for less boilerplate 
            return true
        }

        seen[val] = struct{}{}
    }

    return false
}

