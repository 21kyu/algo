func longestIncreasingPath(matrix [][]int) int {
    h := len(matrix)
    w := len(matrix[0])
    
    if h == 0 {
        return 0
    }
    
    res := 0
    cache := make([][]int, h)
    for i:=0; i<h; i++ {
        cache[i] = make([]int, w)
    }
    
    for i:=0; i<h; i++ {
        for j:=0; j<w; j++ {
            res = max(res, dfs(matrix, i, j, w, h, cache))
        }
    }
    
    return res
}

func dfs(matrix [][]int, i, j, w, h int, cache [][]int) int {
    if cache[i][j] != 0 {
        return cache[i][j]
    }
    
    directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    
    for _, d := range directions {
        x := i + d[0]
        y := j + d[1]
        
        if x >= 0 && y >= 0 && x < h && y < w && (matrix[x][y] > matrix[i][j]) {
            cache[i][j] = max(cache[i][j], dfs(matrix, x, y, w, h, cache))
        }
    }
    
    cache[i][j]++
    return cache[i][j]
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}