/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MinHeap []*ListNode

func (m MinHeap) Len() int { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i].Val < m[j].Val }
func (m MinHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}

func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	item := old[n-1]
	*m = old[:n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
    h := &MinHeap{}
	heap.Init(h)

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		curr.Next = node
		curr = curr.Next

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}
