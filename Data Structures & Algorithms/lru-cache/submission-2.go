type LRUCache struct {
    capacity int
	cache map[int]*list.Element
	list *list.List
}

type Pair struct {
	Key, Value int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		capacity: capacity,
		cache: make(map[int]*list.Element),
		list: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
    if elem, found := this.cache[key]; found {
		this.list.MoveToFront(elem)
		return elem.Value.(*Pair).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if elem, found := this.cache[key]; found {
		elem.Value.(*Pair).Value = value
		this.list.MoveToFront(elem)
		return
	}
	
	if this.list.Len() == this.capacity {
			lruElem := this.list.Back()
			lruPair := lruElem.Value.(*Pair)
			delete(this.cache, lruPair.Key)
			this.list.Remove(lruElem)
	}
	newElem := this.list.PushFront(&Pair{key, value})
	this.cache[key] = newElem
}
