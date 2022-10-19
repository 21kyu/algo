func topKFrequent(words []string, k int) []string {
    h := &MaxHeap{}
    heap.Init(h)
    f := make(map[string]int)
    
    for _, w := range words {
        if _, ok := f[w]; !ok {
            f[w] = 1
        } else {
            f[w]++
        }
    }
    
    for k, v := range f {
        heap.Push(h, Node{k, v})
    }
    
    var ans []string
    for i:=0; i<k; i++ {
        ans = append(ans, heap.Pop(h).(Node).word)
    }
    
    return ans
}

type Node struct {
    word string
    count int
}

type MaxHeap []Node

// len less swap push pop
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
    if h[i].count == h[j].count {
        return h[i].word < h[j].word
    }
    return h[i].count > h[j].count
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(Node))
}
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := old.Len()
    x := old[n-1]
    *h = old[:n-1]
    return x
}