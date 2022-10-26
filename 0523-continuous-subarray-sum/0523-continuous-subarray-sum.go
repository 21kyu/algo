func checkSubarraySum(nums []int, k int) bool {
    m := make(map[int]int, len(nums) + 1)
    m[0] = -1
    rem := 0
    
    for i, n := range nums {
        rem = (rem + n) % k
        if v, ok := m[rem]; ok {
            if i - v >= 2 {
                return true
            }
        } else {
            m[rem] = i    
        }    
    }
    return false
}