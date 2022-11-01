func findBall(grid [][]int) []int {
    m, n := len(grid), len(grid[0])
    var q []Ball
    
    ans := make([]int, n)
    for i:=0; i<n; i++ {
        ans[i] = -1
        q = append(q, Ball{i, 0, i})
    }
    
    for len(q) > 0 {
        b := q[0]
        q = q[1:]
        
        rr, cc := b.r + 1, 0
        
        if grid[b.r][b.c] == 1 {
            if b.c+1 >= n || grid[b.r][b.c+1] != 1 {
                continue
            }
            cc = b.c + 1
        } else {
            if b.c-1 < 0 || grid[b.r][b.c-1] != -1 {
                continue
            }
            cc = b.c - 1
        }
        
        if rr == m {
            ans[b.s] = cc
            continue
        }
        
        q = append(q, Ball{b.s, rr, cc})
    }
    
    return ans
}

type Ball struct {
    s, r, c int
}