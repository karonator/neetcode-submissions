type MyHashSet struct {
	Presence []bool
}

func Constructor() MyHashSet {
	hashset := MyHashSet{}
	hashset.Presence = make([]bool, 1000001)
	return hashset
}

func (this *MyHashSet) Add(key int) {
    this.Presence[key] = true
}

func (this *MyHashSet) Remove(key int) {
    this.Presence[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
    return this.Presence[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 