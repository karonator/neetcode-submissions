type MinStack struct {
	Values 	[]int
	Mins	[]int
}

func Constructor() MinStack {
	return MinStack {
		Values: []int{},
		Mins:	[]int{},
	}
}

func (this *MinStack) Push(val int) {
	this.Values = append(this.Values, val)
	if len(this.Mins) == 0 {
		this.Mins = append(this.Mins, val)
	} else {
		last_min := this.Mins[len(this.Mins) - 1]
		if last_min > val {
			this.Mins = append(this.Mins, val)
		} else {
			this.Mins = append(this.Mins, last_min)
		}
	}
}

func (this *MinStack) Pop() {
	this.Values = this.Values[:len(this.Values) - 1]
	this.Mins = this.Mins[:len(this.Mins) - 1]
}

func (this *MinStack) Top() int {
	return this.Values[len(this.Values) - 1]
}

func (this *MinStack) GetMin() int {
	return this.Mins[len(this.Values) - 1]
}
