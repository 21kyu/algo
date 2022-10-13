/**
 * Definition for an Interval.
 * type Interval struct {
 *     Start int
 *     End   int
 * }
 */

func employeeFreeTime(schedule [][]*Interval) []*Interval {
    var ans []*Interval
    
    if len(schedule) == 0 {
        return ans
    }
    
    // sort!
    // e1 [1,2],[5,6]
    // e2 [1,3]
    // e3 [4,10]
    // 하나의 배열에 정렬 -> [1,2],[1,3],[4,10],[5,6]
    var arr []*Interval
    // O(n)
    for _, s := range schedule {
        for _, e := range s {
            arr = append(arr, e)
        }
    }
    // O(nlogn)
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].Start < arr[j].Start
    })
    
    // 하나의 배열을 순회하면서 (end time을 기록하며) interval이 존재하는지 확인
    endTime := 0
    for i:=0; i<len(arr)-1; i++ {
        if endTime < arr[i].End {
            endTime = arr[i].End
        }
        
        if endTime < arr[i+1].Start {
            interval := &Interval{
                endTime,
                arr[i+1].Start,
            }
            ans = append(ans, interval)
        }
    }
    
    return ans
}