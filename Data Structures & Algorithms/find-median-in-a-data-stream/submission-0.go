type MedianFinder struct {
    data []int
}


func Constructor() MedianFinder {
    return MedianFinder{}
}


func (this *MedianFinder) AddNum(num int)  {
    this.data = append(this.data, num)
}


func (this *MedianFinder) FindMedian() float64 {
    sort.Ints(this.data)
    n := len(this.data)
    if n % 2 != 0 {
        return float64(this.data[n/2])
    }
    return float64(this.data[n/2] + this.data[n/2-1]) / 2.0 
}
