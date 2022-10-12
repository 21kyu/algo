func largestPerimeter(nums []int) int {
    // a <= b <= c
    sort.Ints(nums)
    
    for i:=len(nums)-3; i>=0; i-- {
        // a + b > c
        if nums[i] + nums[i+1] > nums[i+2] {
            return nums[i] + nums[i+1] + nums[i+2]
        }
    }
    
    return 0
}