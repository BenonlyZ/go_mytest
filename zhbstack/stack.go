package main

import "fmt"

type Element interface{}

type Stack struct {
	list []Element
}

//初始化栈
func NewStack() *Stack {
	return &Stack{
		list: make([]Element, 0),
	}
}

func (s *Stack) Len() int {
	return len(s.list)
}

//判断栈是否空
func (s *Stack) IsEmpty() bool {
	if len(s.list) == 0 {
		return true
	} else {
		return false
	}
}

//入栈
func (s *Stack) Push(x interface{}) {
	s.list = append(s.list, x)
}

//连续传入
func (s *Stack) PushList(x []Element) {
	s.list = append(s.list, x...)
}

//出栈
func (s *Stack) Pop() Element {
	if len(s.list) <= 0 {
		fmt.Println("Stack id Empty")
		return nil
	} else {
		ret := s.list[len(s.list)-1]
		s.list = s.list[:len(s.list)-1]
		return ret
	}
}

//返回栈顶元素
func (s *Stack) Top() Element {
	if s.IsEmpty() == true {
		fmt.Println("Stack is Empty")
		return nil
	} else {
		return s.list[len(s.list)-1]
	}
}

//清空栈
func (s *Stack) Clear() bool {
	if len(s.list) == 0 {
		return true
	}
	for i := 0; i < s.Len(); i++ {
		s.list[i] = nil
	}
	s.list = make([]Element, 0)
	return true
}

//打印测试
func (s *Stack) Show() {
	len := len(s.list)
	for i := 0; i != len; i++ {
		fmt.Println(s.Pop())
	}
}

//交换值
func (stack *Stack) Swap(other *Stack) {
	switch {
	case len(stack.list) == 0 && len(other.list) == 0:
		return
	case len(other.list) == 0:
		other.list = stack.list[:len(stack.list)]
		stack.list = nil
	case len(stack.list) == 0:
		stack.list = other.list
		other.list = nil
	default:
		stack.list, other.list = other.list, stack.list
	}
	return
}

func main001() {
	stack := NewStack()
	stack2 := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	//	stack.Pop()
	stack2.PushList([]Element{4, 5, 6, 7})
	stack.Swap(stack2)
	fmt.Println("栈1如下：")
	stack.Show()
	fmt.Println("栈2如下：")
	stack2.Show()
	//	fmt.Println("Stack len is ", stack.Len())
	//	fmt.Println("Top is", stack.Top()) //栈顶元素
	//stack.Show()                       //show之后数据已经清空
}

//判断括号(,{,[字符串是否有效
func isValid(s string) bool {
	var str = make([]byte, 0)
	count := len(str)
	for i := 0; i < len(s); i++ {

		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			str = append(str, s[i])
			count++
		} else {
			if count == 0 {
				return false
			}
			switch s[i] {
			case ')':
				if str[count-1] == '(' {
					str = str[0 : count-1]
					count--
				} else {
					return false
				}
			case ']':
				if str[count-1] == '[' {
					str = str[0 : count-1]
					count--
				} else {
					return false
				}
			case '}':
				if str[count-1] == '{' {
					str = str[0 : count-1]
					count--
				} else {
					return false
				}
			}
		}
	}
	if count == 0 {
		return true
	} else {
		return false
	}
}

//整数反转
func reverse(x int) int {
	if x == 0 {
		return 0
	}
	res := 0
	for x != 0 {
		y := x % 10
		x = x / 10
		res = res*10 + y
	}
	return res
}
