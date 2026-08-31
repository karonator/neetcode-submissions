const storageSize = 1017;

type MyHashSet struct {
	data [][]int
}

func Constructor() MyHashSet {
    return MyHashSet{
		data: make([][]int, storageSize),
	}
}

func (this *MyHashSet) hashInt(key int) int {
    return key % storageSize
}


func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}

    idx := this.hashInt(key)
	if this.data[idx] == nil {
		this.data[idx] = []int{key}
	} else {
		this.data[idx] = append(this.data[idx], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	if !this.Contains(key) {
		return
	}

    idx := this.hashInt(key)
    for i, n := range(this.data[idx]) {
		if n == key {
			this.data[idx][len(this.data[idx]) - 1], this.data[idx][i] = this.data[idx][i], this.data[idx][len(this.data[idx]) - 1]
			break
		}
	}
	this.data[idx] = this.data[idx][:len(this.data[idx]) - 1]
}

func (this *MyHashSet) Contains(key int) bool {
    idx := this.hashInt(key)
	if this.data[idx] == nil {
		return false
	}
	for _, n := range(this.data[idx]) {
		if n == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 