func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    w1, w2 := "", ""
    
    for _, str := range word1 {
        w1 += str
    }
    
    for _, str := range word2 {
        w2 += str
    }
    
    return w1 == w2
}