func shortestPath(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    h := &MinHeap{}
    heap.Init(h)
    heap.Push(h, Item{0, 0, 0, k+1})
    dir := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    visited := make([][]int, m)
    for i := range visited {
        visited[i] = make([]int, n)
    }
    visited[0][0] = k+1
    
    for h.Len() > 0 {
        item := heap.Pop(h).(Item)
        r, c, s := item.row, item.col, item.step
        
        if r == m-1 && c == n-1 {
            return s
        }
        
        for _, d := range dir {
            rr, cc, e := r + d[0], c + d[1], item.eliminateCount
            
            if rr >= 0 && cc >= 0 && rr < m && cc < n {
                if grid[rr][cc] == 1 {
                    e--
                    if e == 0 {
                        continue
                    }
                }
                
                if visited[rr][cc] < e {
                    visited[rr][cc] = e
                    heap.Push(h, Item{rr, cc, s+1, e})
                }
            }
        }
    }
    
    return -1
}

type Item struct {
    row, col int
    step, eliminateCount int
}

type MinHeap []Item

// Len Less Swap Push Pop
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].step < h[j].step }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Item))
}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := old.Len()
    x := old[n-1]
    *h = old[:n-1]
    return x
}