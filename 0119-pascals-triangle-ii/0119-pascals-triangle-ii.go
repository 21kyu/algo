func getRow(rowIndex int) []int {
    pascal := []int{1}
    
    var helper func(int, []int) []int
    helper = func(i int, p []int) []int {
        if i == rowIndex {
            return p
        }
        
        p = append(p, 1)
        for i:=len(p)-2; i>0; i-- {
            p[i] += p[i-1]
        }
        
        return helper(i+1, p)
    }
    
    return helper(0, pascal)
    
}

