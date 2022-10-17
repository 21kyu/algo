func checkIfPangram(sentence string) bool {
    set := make(map[byte]bool)
    
    for i := range sentence {
        set[sentence[i]] = true
        if len(set) == 26 {
            return true
        }
    }
    
    return len(set) == 26
}