package main

type MyStack struct {
	queue1 []int
	queue2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue1 = append(this.queue1, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	var re int
	var i int
	l1 := this.queue1
	l2 := this.queue2
	switch {
	case len(l1) == 0 && len(l2) == 0:
		re = -1
	case len(l1) == 0:
		for i = 0; i < len(l2)-1; i++ {
			this.queue1 = append(this.queue1, this.queue2[i])
		}
		re = this.queue2[i]
		this.queue2 = nil
	case len(l2) == 0:
		for i = 0; i < len(l1)-1; i++ {
			this.queue2 = append(this.queue2, this.queue1[i])
		}
		re = this.queue1[i]
		this.queue1 = nil
	}
	return re
}

/** Get the top element. */
func (this *MyStack) Top() int {
	var re int
	switch {
	case len(this.queue1) == 0:
		re = this.queue2[len(this.queue2)-1]
	case len(this.queue2) == 0:
		re = this.queue1[len(this.queue1)-1]
	}
	return re
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.queue1) == 0 && len(this.queue2) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
