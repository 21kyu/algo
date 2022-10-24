func maxLength(arr []string) int {
    n := len(arr)
    
    var backTrack func(int, string) int
    backTrack = func(index int, form string) int {
        var m [26]bool
        for i := range form {
            if m[form[i]-'a'] {
                return 0
            }
            m[form[i]-'a'] = true
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