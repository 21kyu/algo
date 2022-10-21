func containsNearbyDuplicate(nums []int, k int) bool {
    m := make(map[int]int)
    
    for i := range nums {
        if v, ok := m[nums[i]]; ok {
            if abs(i - v) <= k {
                return true
            }
        }
        m[nums[i]] = i
    }
    
    return false
}

func abs(x int) int {
    if x < 0 {
        return x * -1
    }
    return x
}