type HashItem struct {
	Key int
	Val int
}

type MyHashMap struct {
	StorageSize int
	Data [][]*HashItem
}

func Constructor() MyHashMap {
	size := 10000
    return MyHashMap {
		StorageSize: size,
		Data: make([][]*HashItem, size),
	}
}

func (this *MyHashMap) Put(key int, value int) {
    idx := key % this.StorageSize
	found := false
	for _, item := range this.Data[idx] {
		if item.Key == key {
			found = true
			item.Val = value
			break
		}
	}
	if !found {
		this.Data[idx] = append(this.Data[idx], &HashItem {
			Key: key,
			Val: value,
		})
	}
}

func (this *MyHashMap) Get(key int) int {
    idx := key % this.StorageSize
	for _, item := range this.Data[idx] {
		if item.Key == key {
			return item.Val
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	idx := key % this.StorageSize
	for i, item := range this.Data[idx] {
		if item.Key == key {
			this.Data[idx][i] = this.Data[idx][len(this.Data[idx]) - 1]
			this.Data[idx] = this.Data[idx][:len(this.Data[idx]) - 1]
			break
		}
	}

}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */