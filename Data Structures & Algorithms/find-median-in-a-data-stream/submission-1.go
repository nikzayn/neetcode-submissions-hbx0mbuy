type MinHeap []int

func (m MinHeap) Len() int { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m MinHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m MinHeap) Peek() int { return m[0] }

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

type MaxHeap []int

func (m MaxHeap) Len() int { return len(m) }
func (m MaxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m MaxHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m MaxHeap) Peek() int { return m[0] }

func (m *MaxHeap) Push(x interface{}) {
    *m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
    old := *m
    n := len(old)
    x := old[n-1]
    *m = old[:n-1]
    return x
}


type MedianFinder struct {
    lower *MaxHeap
    upper *MinHeap
}


func Constructor() MedianFinder {
    l := &MaxHeap{}
    u := &MinHeap{}

    heap.Init(l)
    heap.Init(u)

    return MedianFinder{
        lower: l,
        upper: u,
    }
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.lower, num)
    heap.Push(this.upper, heap.Pop(this.lower).(int))

    if this.upper.Len() > this.lower.Len() {
        heap.Push(this.lower, heap.Pop(this.upper).(int))
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.lower.Len() > this.upper.Len() {
        return float64(this.lower.Peek())
    }
    return float64(this.lower.Peek() + this.upper.Peek()) / 2.0
}
