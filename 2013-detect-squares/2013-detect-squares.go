type DetectSquares struct {
    counts map[[2]int]int
    points [][]int
}


func Constructor() DetectSquares {
    return DetectSquares{
        make(map[[2]int]int),
        [][]int{},
    }
}


func (d *DetectSquares) Add(point []int)  {
    d.counts[[2]int{point[0], point[1]}]++
    d.points = append(d.points, point)
}


func (d *DetectSquares) Count(point []int) int {
    x, y := point[0], point[1]
    ans := 0
    
    for _, p := range d.points {
        if abs(x - p[0]) != abs(y - p[1]) || x == p[0] || y == p[1] {
            continue
        }
        
        ans += d.counts[[2]int{x, p[1]}] * d.counts[[2]int{p[0], y}]
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return x * -1
    }
    return x
}


/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */