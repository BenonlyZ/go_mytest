package main

import "fmt"

type stackNode struct {
	data interface{}
	next *stackNode
}

type stackList struct {
	length   int
	headNode *stackNode
}

//初始化栈
func initLinkStack() *stackList {
	node := new(stackNode)
	L := new(stackList)
	L.headNode = node
	L.length = 0
	return L
}

//入栈
func (stack *stackList) push(val interface{}) {
	node := new(stackNode)
	node.data = val
	node.next = stack.headNode
	stack.headNode = node
	stack.length++
}

//出栈
func (stack *stackList) pop() interface{} {
	if stack.headNode == nil {
		return nil
	}
	val := stack.headNode.data
	stack.headNode = stack.headNode.next
	stack.length--
	return val
}

//取栈顶元素
func (stack *stackList) getTop() interface{} {
	if stack.headNode == nil {
		return nil
	}
	val := stack.headNode.data
	return val
}

//查看栈内所有元素
func (stack *stackList) showAll() {
	if stack.headNode.next == nil {
		fmt.Println("空栈")
		return
	}
	cur := stack.headNode
	for {
		if cur.next != nil {
			fmt.Printf("%v\n", cur.data)
			cur = cur.next
		} else {
			break
		}
	}
}

func main() {
	L := initLinkStack()
	arr := []interface{}{
		"aa",
		"bb",
		"cc",
		"dd",
	}
	for i := range arr {
		L.push(arr[i])
	}
	L.showAll()
	fmt.Println(L.getTop())
	fmt.Println(L.pop())
	fmt.Println(L.getTop())
}
