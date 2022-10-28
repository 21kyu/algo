func groupAnagrams(strs []string) [][]string {
    
    // key: a 1, e 1, t 1
    // value : eat
    m := make(map[[26]int][]string)
    
    for _, s := range strs {
        var chAry [26]int
        
        for i := range s {
            chAry[s[i]-'a']++
        }
        
        m[chAry] = append(m[chAry], s)
    }
    
    var ans [][]string
    
    for _, item := range m {
        ans = append(ans, item)
    }
    
    return ans
}