package main

import "fmt"

type MinStack struct {
	stack []item
}

type item struct {
	x   int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{min: min, x: x})
}

func (this *MinStack) Pop() {

	this.stack = this.stack[0 : len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].x
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}
func main() {
	obj := Constructor()
	obj.Push(5)
	obj.Push(4)
	obj.Push(4)
	obj.Push(6)
	obj.Push(2)
	fmt.Println(obj)
	fmt.Println(obj.Top)
	obj.Pop()
	obj.Pop()
	fmt.Println(obj)

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
