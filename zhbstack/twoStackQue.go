package main

//两个栈实现一个队列
type MyQueue struct {
	stack1 []int
	stack2 []int
}

//初始化这个队列
func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

//入队列：选择队列里的stack1栈，进行入栈；结束后则stack1栈有数据，stack2栈无数据
func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

//出队列：由于stack2栈有元素，则说明已经经过stack1->stack2的操作，则直接对stack2出栈即可；如果stack2为空，则先进行stack1->stack2的操作
func (this *MyQueue) Pop() int {
	if len(this.stack2) == 0 {
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
			this.stack1 = this.stack1[:i]
		}
	}
	re := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return re
}

//同上，只是不对stack2进行出栈操作，只是返回stack2的栈顶
func (this *MyQueue) Peek() int {
	if len(this.stack2) == 0 {
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
			this.stack1 = this.stack1[:i]
		}
	}
	re := this.stack2[len(this.stack2)-1]
	return re
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
