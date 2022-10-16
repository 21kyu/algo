// 1. dijkstra + min heap
// 2. MST, kruskal + unionfind + sort
// 3. binary search..??

func minimumEffortPath(heights [][]int) int {
    // return minimalSpanningTree(heights)
    // return dijkstraWithMinHeap(heights)
    return binarySearchAndDfs(heights)
}

func binarySearchAndDfs(heights [][]int) int {
    n, m := len(heights), len(heights[0])
    low, high := 0, 0
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            high = max(high, heights[i][j])
        }
    }
    ans := high
    
    reachable := func(mid int) bool {
        // i, j -> i*m + j
        seen := make(map[int]bool)
        stk := []int{0}
        dir := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
        seen[0] = true
        
        for len(stk) > 0 {
            x := stk[len(stk)-1]
            stk = stk[:len(stk)-1]
            r, c := x/m, x%m
            
            if r == n-1 && c == m-1 {
                return true
            }
            
            for _, d := range dir {
                rr, cc := r + d[0], c + d[1]
                xx := rr * m + cc
                if rr>=0 && cc>=0 && rr<n && cc<m && !seen[xx] && abs(heights[rr][cc] - heights[r][c]) <= mid {
                    stk = append(stk, xx)
                    seen[xx] = true
                }
            }
        }
        
        return false
    }
    
    for low <= high {
        mid := int(uint(low+high) >> 1)
        if reachable(mid) {
            ans = min(ans, mid)
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    
    return ans
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

// vlogv (v = n*m)
func dijkstraWithMinHeap(heights [][]int) int {
    // min heap을 만든다
    n, m := len(heights), len(heights[0])
    h := &MinHeap{}
    heap.Init(h)
    // 시작 0,0 을 push
    heap.Push(h, Item{0, 0, 0})
    // visit 배열
    visited := make([][]bool, n)
    for i := range visited {
        visited[i] = make([]bool, m)
    }
    
    dir := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    
    // min heap 기준으로 돌면서
    for h.Len() > 0 {
        item := heap.Pop(h).(Item)
        r, c, v := item.r, item.c, item.v
        
        visited[r][c] = true
        
        // 맨 끝에 도달했으면 diff 값 반환
        if r == n-1 && c == m-1 {
            return v
        }
        
        // 4 방향 확인
        for _, d := range dir {
            rr, cc := r + d[0], c + d[1]
            // 아직 방문하지 않았다면 크기 계산해서 heap에 push
            if rr >= 0 && cc >= 0 && rr < n && cc < m && !visited[rr][cc] {
                heap.Push(h, Item{max(v, abs(heights[rr][cc] - heights[r][c])), rr, cc})
            }
        }
    }
    
    return -1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
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