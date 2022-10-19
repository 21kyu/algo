func fullJustify(words []string, maxWidth int) []string {
    var result []string
    var tokens []string
    
    checkCount := 0
    realCount := 0
    
    for _, w := range words {
        l := len(w)
        
        if checkCount + l <= maxWidth {
            tokens = append(tokens, w)
            checkCount += l + 1
            realCount += l
        } else {
            result = append(result, padExtraSpaces(tokens, maxWidth, realCount, len(words) - 1 == l))
            tokens = []string{w}
            checkCount = l + 1
            realCount = l
        }
    }
    
    if checkCount > 0 {
        result = append(result, padExtraSpaces(tokens, maxWidth, realCount, true))
    }
    
    return result
}

func padExtraSpaces(tokens []string, maxWidth int, count int, last bool) string {
    l := len(tokens)
    
    if l == 1 {
        return tokens[0] + strings.Repeat(" ", maxWidth - count)
    } else if last {
        return strings.Join(tokens, " ") + strings.Repeat(" ", maxWidth - count - (l - 1))
    } else {
        slots := l - 1
        spaces := maxWidth - count
        quotient := spaces / slots
        remains := spaces % slots
        
        for i:=0; i<slots; i++ {
            if remains > 0 {
                tokens[i] = tokens[i] + strings.Repeat(" ", quotient + 1)
            } else {
                tokens[i] = tokens[i] + strings.Repeat(" ", quotient)
            }
            remains--
        }
        
        return strings.Join(tokens, "")
    }
}