/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    first, second := head, head.Next
    first.Next, second.Next = swapPairs(second.Next), first
    return second
}
