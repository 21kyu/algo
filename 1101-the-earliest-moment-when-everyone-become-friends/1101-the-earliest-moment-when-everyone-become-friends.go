func earliestAcq(logs [][]int, n int) int {
    return unionFind(logs, n)
}

func unionFind(logs [][]int, n int) int {
    if len(logs) < n-1 {
        return -1
    }
    
    // O(n)
    ds := NewDisjointSet(n)
    
    // O(nlogn)
    sort.Slice(logs, func(a, b int) bool {
        return logs[a][0] < logs[b][0]
    })
    
    // O(a(N)) => O(4)
    for i := range logs {
        if ds.union(logs[i][1], logs[i][2]) {
            n--
        }
        
        if n == 1 {
            return logs[i][0]
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
    if ds.parent[x] != x {
        ds.parent[x] = ds.find(ds.parent[x])
    }
    
    return ds.parent[x]
}

func (ds *DisjointSet) union(x, y int) bool {
    xRoot, yRoot := ds.find(x), ds.find(y)
    
    if xRoot == yRoot {
        return false
    }
    
    // union by rank
    if ds.rank[xRoot] > ds.rank[yRoot] {
        ds.parent[yRoot] = xRoot
    } else if ds.rank[yRoot] > ds.rank[xRoot] {
        ds.parent[xRoot] = yRoot
    } else {
        ds.parent[xRoot] = yRoot
        // path compression
        ds.rank[xRoot]++
    }
    
    return true
}