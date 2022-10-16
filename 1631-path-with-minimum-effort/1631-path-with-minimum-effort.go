// 1. dijkstra + min heap
// 2. MST, kruskal + unionfind
// 3. binary search..??

func minimumEffortPath(heights [][]int) int {
    return dijkstraWithMinHeap(heights)
}

func dijkstraWithMinHeap(heights [][]int) int {
    n, m := len(heights), len(heights[0])
    h := &MinHeap{}
    heap.Init(h)
    heap.Push(h, Item{0, 0, 0})
    
    visited := make([][]bool, n)
    diff := make([][]int, n)
    for i := range diff {
        visited[i] = make([]bool, m)
        diff[i] = make([]int, m)
    }
    
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            diff[i][j] = math.MaxInt
        }
    }
    
    visited[0][0] = true
    diff[0][0] = 0
    
    dir := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    
    for h.Len() > 0 {
        item := heap.Pop(h).(Item)
        r, c := item.r, item.c
        
        visited[r][c] = true
        diff[r][c] = item.v
        
        if r == n-1 && c == m-1 {
            return item.v
        }
        
        for _, d := range dir {
            rr, cc := r + d[0], c + d[1]
            if rr >= 0 && cc >= 0 && rr < n && cc < m  {
                effort := abs(heights[rr][cc] - heights[r][c])
                if effort < item.v {
                    effort = item.v
                }
                if diff[rr][cc] > effort {
                    heap.Push(h, Item{effort, rr, cc})
                }
            }
        }
    }
    
    return -1
}

func abs(x int) int {
    if x < 0 {
        return x * -1
    }
    return x
}

type Item struct {
    v, r, c int
}

type MinHeap []Item

// Len Less Swap Push Pop
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].v < h[j].v }
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