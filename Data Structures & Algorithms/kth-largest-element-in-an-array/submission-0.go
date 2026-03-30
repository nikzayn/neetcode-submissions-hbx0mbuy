type MinHeap []int

func (m MinHeap) Len() int {return len(m)}
func (m MinHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m MinHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(x interface{}) {
    *m = append(*m, x.(int))
}

func (m *MinHeap) Pop() interface{} {
    old := *m
    n := len(old)
    x := old[n-1]
    *m = old[:n-1]
    return x
}

func findKthLargest(nums []int, k int) int {
    h := &MinHeap{}
    heap.Init(h)

    for _, num := range nums {
       heap.Push(h, num)
       if h.Len() > k {
            heap.Pop(h)
       }
    }

    return (*h)[0]
}
