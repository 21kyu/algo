func maxLength(arr []string) int {
    n := len(arr)
    
    var backTrack func(int, string) int
    backTrack = func(index int, form string) int {
        m := make(map[rune]bool, len(form))
        for _, ch := range form {
            if _, ok := m[ch]; ok {
                return 0
            }
            m[ch] = true
        }
        
        if index >= n {
            return len(form)
        }
        
        return max(backTrack(index+1, form), backTrack(index+1, form+arr[index]))
    }
    
    return backTrack(0, "")
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}