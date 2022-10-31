func reverseString(s []byte)  {
    
    var helper func([]byte, int, int)
    helper = func(s []byte, l, r int) {
        s[l], s[r] = s[r], s[l]
        if l+1 < r {
            helper(s, l+1, r-1)
        }
    }
    
    helper(s, 0, len(s)-1)
}
