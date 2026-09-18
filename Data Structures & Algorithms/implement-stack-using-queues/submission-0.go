type MyStack struct {
	data []int
}

func Constructor() MyStack {
	return MyStack{
		data: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < len(this.data) - 1; i++ {
		first := this.data[0]
		this.data = this.data[1: len(this.data)]
		this.data = append(this.data, first)
	}
	first := this.data[0]
	this.data = this.data[1: len(this.data)]
	return first
}

func (this *MyStack) Top() int {
	ans := 0
	for i := 0; i < len(this.data); i++ {
		first := this.data[0]
		this.data = this.data[1: len(this.data)]
		this.data = append(this.data, first)
		if i == len(this.data) - 1 {
			ans = first
		}
	}
	return ans
}

func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
