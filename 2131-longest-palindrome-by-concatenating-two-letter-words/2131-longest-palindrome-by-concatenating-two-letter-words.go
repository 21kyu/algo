func longestPalindrome(words []string) int {
    seen := make(map[string]int)
    length := 0
    
    for _, w := range words {
        r := string([]byte{w[1], w[0]})
        
        if i, ok := seen[r]; ok && i > 0 {
            seen[r]--
            length += 4
        } else {
            if _, ok := seen[w]; ok {
                seen[w]++
            } else {
                seen[w] = 1
            }
        }
    }
    
    for k, v := range seen {
        if k[0] == k[1] && v > 0 {
            length += 2
            break
        }
    }
    
    return length
}
