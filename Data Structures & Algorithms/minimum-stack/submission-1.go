type MinStack struct {
	data []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		data: []int{},
		mins: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, min(val, this.mins[len(this.mins) - 1]))
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data) - 1]
	this.mins = this.mins[:len(this.mins) - 1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data) - 1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins) - 1]
}
