func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    return hash(word1) == hash(word2)
}

func hash(word []string) int {
    hash := 5381
    for _, str := range word {
        for _, ch := range str {
            hash = (hash * 33) ^ int(ch)
        }
    }
    return hash
}