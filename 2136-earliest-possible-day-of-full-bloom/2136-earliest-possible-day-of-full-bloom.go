func earliestFullBloom(plantTime []int, growTime []int) int {
    n := len(growTime)
    indices := make([]int, n)
    
    for i := range indices {
        indices[i] = i
    }
    
    sort.Slice(indices, func(i, j int) bool {
        return growTime[indices[i]] > growTime[indices[j]]
    })
    
    plant, maxGrow := 0, 0
    
    for _, index := range indices {
        plant += plantTime[index]
        maxGrow = max(maxGrow, plant + growTime[index])
    }
    
    return maxGrow
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}