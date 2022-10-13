var memo [1001]int

func getKth(lo int, hi int, k int) int {
    var res [][]int
    
    for i:=lo; i<=hi; i++ {
        res = append(res, []int{i, recur(i)})
    }
    
    sort.Slice(res, func(a, b int) bool {
        return res[a][1] < res[b][1] || (res[a][1] == res[b][1] && res[a][0] < res[b][0])
    })
    
    return res[k-1][0]
}

func recur(n int) int {
    if n == 1 {
        return 0
    }
    
    if n < 1001 && memo[n] != 0 {
        return memo[n]
    }

    var next int
    if n%2 == 0 {
        next = n/2
    } else {
        next = 3*n + 1
    }

    ans := 1 + recur(next)
    
    if n < 1001 {
        memo[n] = ans
    }
    return ans
}