type LRUCache struct {
    cache map[int]*list.Element
	capacity int
	list *list.List
}

type Pair struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		cache: make(map[int]*list.Element),
		capacity: capacity,
		list: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
    if elem, found := this.cache[key]; found {
		this.list.MoveToFront(elem)
		return elem.Value.(*Pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if elem, found := this.cache[key]; found {
		elem.Value.(*Pair).value = value
		this.list.MoveToFront(elem)
		return
	}

	if this.list.Len() == this.capacity {
		lruElem := this.list.Back()
		lruPair := lruElem.Value.(*Pair)
		delete(this.cache, lruPair.key)
		this.list.Remove(lruElem)
	}
	newElem := this.list.PushFront(&Pair{key, value})
	this.cache[key] = newElem
}
