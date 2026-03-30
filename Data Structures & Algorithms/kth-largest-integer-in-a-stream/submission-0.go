type MinHeap []int

func (m MinHeap) Len() int { return len(m) }
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

type KthLargest struct {
    minHeap *MinHeap
    k int
}


func Constructor(k int, nums []int) KthLargest {
    h := &MinHeap{}
    heap.Init(h)

    k1 := KthLargest {
        minHeap: h,
        k: k,
    }

    for _, num := range nums {
        k1.Add(num)
    }

    return k1
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.minHeap, val)
    if this.minHeap.Len() > this.k {
        heap.Pop(this.minHeap)
    }
    return (*this.minHeap)[0]
}
