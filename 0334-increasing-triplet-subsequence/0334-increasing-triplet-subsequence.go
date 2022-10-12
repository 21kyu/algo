func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {
        return false
    }
    
    f, s := math.MaxInt, math.MaxInt
    
    // 2 1 5 0 4 6
    // f f s f s *
    for _, n := range nums {
        if f >= n {
            f = n
        } else if s >= n {
            s = n
        } else {
            return true
        }
    }
    
    return false
}