func swimInWater(grid [][]int) int {
    return dijkstraWithMinHeap(grid)
}

func dijkstraWithMinHeap(grid [][]int) int {
    q := &MinHeap{}
    heap.Init(q)
    heap.Push(q, Item{grid[0][0], 0, 0})
    
    n := len(grid)
    visited := make([][]bool, n)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    
    dir := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
    
    // 0,0 -> n-1,n-1까지 방문하면서 max(prev.h, cur.h) 값을 h에 유지해준다
    for q.Len() > 0 {
        x := heap.Pop(q).(Item)
        
        if x.r == n-1 && x.c == n-1 {
            return x.h
        }
        
        visited[x.r][x.c] = true
        
        for _, d := range dir {
            rr, cc := x.r + d[0], x.c + d[1]
            if rr >= 0 && cc >= 0 && rr < n && cc < n && !visited[rr][cc] {
                heap.Push(q, Item{max(x.h, grid[rr][cc]), rr, cc})
            }
        }
    }
    
    return -1
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}

type Item struct {
    h, r, c int
}

type MinHeap []Item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].h < h[j].h }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Item))
}
func (h *MinHeap) Pop() interface{} {
    tmp := *h
    n := tmp.Len()
    x := tmp[n-1]
    *h = tmp[:n-1]
    return x
}