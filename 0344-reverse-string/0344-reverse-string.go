func reverseString(s []byte)  {
    helper(s, 0, len(s)-1)
}

func helper(s []byte, l, r int) {
    if l >= r {
        return
    }
    
    s[l], s[r] = s[r], s[l]
    helper(s, l+1, r-1)
}