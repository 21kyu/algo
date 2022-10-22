func minWindow(s string, t string) string {
    m, n := len(s), len(t)
    
    if m < n {
        return ""
    }
    
    rem := 0
    hashMap := make(map[byte]int)
    
    // ABC -> map{A:1, B:1, C:1} / rem=3
    for i := range t {
        hashMap[t[i]]++
        rem++
    }
    
    isUpdated := false
    ans := string(make([]byte, m))
    left, right := 0, 0
    
    // ADOBECODEBANC two pointer
    // |
    // left & right
    for right < m {
        // right ->
        // right 위치의 값이 hashMap에 있다면 rem, hashMap value -1
        if v, ok := hashMap[s[right]]; ok {
            if v > 0 {
                rem--
            }
            hashMap[s[right]]--
        }
        
        // left ->
        for rem <= 0 {
            // ans 업데이트 가능한지
            if len(ans) >= len(s[left:right+1]) {
                ans = s[left:right+1]
                isUpdated = true
            }
            
            // left 위치의 값이 hashMap에 있다면 rem, hashMap value +1
            if v, ok := hashMap[s[left]]; ok {
                if v > -1 {
                    rem++
                }
                hashMap[s[left]]++
            }
            left++
        }
        right++
    }
    
    if !isUpdated {
        return ""
    }
    
    return ans
}