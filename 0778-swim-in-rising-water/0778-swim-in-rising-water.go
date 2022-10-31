func swimInWater(grid [][]int) int {
    //return byDisjointSetUnion(grid)
    return binarySearchAndDfs(grid)
    //return dijkstraWithMinHeap(grid)
}

// Minimal Spanning Tree / Kruskal
func byDisjointSetUnion(grid [][]int) int {
    n := len(grid)
    
    if n == 1 {
        return 0
    }
    
    ds := NewDisjointSet(n*n)
    pos := make([]int, n*n)
    dir := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    
    for i := range pos {
        pos[i] = i
    }
    
    // n^2*logn^2
    sort.Slice(pos, func(a, b int) bool {
        return grid[pos[a]/n][pos[a]%n] < grid[pos[b]/n][pos[b]%n]
    })
    
    // n^2
    for _, p := range pos {
        r, c := p/n, p%n
        
        for _, d := range dir {
            rr, cc := r + d[0], c + d[1]
            
            if rr >= 0 && cc >= 0 && rr < n && cc < n && grid[rr][cc] < grid[r][c] {
                // a(n^2)
                ds.union(p, rr*n+cc)
            }
        }
        
        if ds.find(0) == ds.find(n*n-1) {
            return grid[r][c]
        }
    }
    
    return -1
    
}

type DisjointSet struct {
    rank, parent []int
    size int
}

func NewDisjointSet(n int) *DisjointSet {
    ds := &DisjointSet{
        make([]int, n),
        make([]int, n),
        n,
    }
    for i:=0; i<n; i++ {
        ds.parent[i] = i
    }
    return ds
}

func (ds *DisjointSet) find(x int) int {
    if x != ds.parent[x] {
        ds.parent[x] = ds.find(ds.parent[x])
    }
    return ds.parent[x]
}

func (ds *DisjointSet) union(x, y int) bool {
    px, py := ds.find(x), ds.find(y)
    
    if px == py {
        return false
    }
    
    if ds.rank[px] > ds.rank[py] {
        ds.parent[py] = px
    } else {
        ds.parent[px] = py
        if ds.rank[px] == ds.rank[py] {
            ds.rank[py]++
        }
    }
    
    return true
}

func binarySearchAndDfs(grid [][]int) int {
    n := len(grid)
    l, h := grid[0][0], n*n
    
    // t보다 작거나 같은 시간에 n-1, n-1까지 도달할 수 있는지를 확인
    possible := func(t int) bool {
        seen := make(map[int]bool)
        seen[0] = true
        dir := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
        
        // 빠르게 끝점까지 도달해보기 위해서 dfs 형태로 탐색
        stk := []int{0}
        
        for len(stk) > 0 {
            k := stk[len(stk)-1]
            stk = stk[:len(stk)-1]
            r, c := k/n, k%n
            
            if r == n-1 && c == n-1 {
                return true
            }
            
            for _, d := range dir {
                rr, cc := r+d[0], c+d[1]
                kk := rr * n + cc
                
                if rr >= 0 && rr < n && cc >= 0 && cc < n && !seen[kk] && grid[rr][cc] <= t {
                    stk = append(stk, kk)
                    seen[kk] = true
                }
            }
        }
        return false
    }
    
    // logn 이진탐색
    for l < h {
        m := int(uint(l+h) >> 1)
        // n^2
        if possible(m) {
            // 현재 높이(m)으로 맵의 끝까지 도달 가능하면 h를 m으로 설정
            h = m
        } else {
            // 도달 불가능하면 l를 m+1로 설정
            l = m + 1
        }
    }
    
    return l
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