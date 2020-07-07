package minStack

import "math"

type MinStack struct {
	stack    []int
	minValue int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0, 0),
		minValue: math.MaxInt32,
	}
}

func (this *MinStack) Push(x int) {
	if x <= this.minValue {
		this.stack = append(this.stack, this.minValue)
		this.minValue = x
	}
	this.stack = append(this.stack, x)
	return
}

func (this *MinStack) Pop() {
	top := len(this.stack) - 1
	topValue := this.stack[top]
	this.stack = this.stack[:top]
	if topValue == this.minValue {
		this.minValue = this.stack[top-1]
		this.stack = this.stack[:top-1]
	}
}

func (this *MinStack) Top() int {
	top := len(this.stack) - 1
	return this.stack[top]
}

func (this *MinStack) GetMin() int {
	return this.minValue
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
