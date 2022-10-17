func checkIfPangram(sentence string) bool {
    alpha := [26]byte{}
    
    for i := range sentence {
        alpha[sentence[i] - 'a']++
    }
    
    for _, v := range alpha {
        if v == 0 {
            return false
        }
    }
    
    return true
}