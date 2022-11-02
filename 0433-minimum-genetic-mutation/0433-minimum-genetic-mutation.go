func minMutation(start string, end string, bank []string) int {
    set := make(map[string]bool)
    sean := make(map[string]bool)
    
    for _, gene := range bank {
        set[gene] = true
    }
    
    var queue []Gene
    queue = append(queue, Gene{start, 0})
    sean[start] = true
    
    choice := []byte{'A', 'C', 'G', 'T'}
    
    for len(queue) > 0 {
        item := queue[0]
        queue = queue[1:]
        
        if item.gene == end {
            return item.numberOfMutations
        }
        
        for _, ch := range choice {
            for i := range item.gene {
                if ch == item.gene[i] {
                    continue
                }
                
                newGene := item.gene[:i] + string(ch) + item.gene[i+1:]
                
                if _, ok := sean[newGene]; ok {
                    continue
                }
                
                sean[newGene] = true
                
                if _, ok := set[newGene]; ok {
                    queue = append(queue, Gene{newGene, item.numberOfMutations + 1})
                }
            }
        }
    }
    
    return -1
}

type Gene struct {
    gene string
    numberOfMutations int
}