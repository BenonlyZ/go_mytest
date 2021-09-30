package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	headNode *Node
}

func main() {
	var ls = new(List) //不随插入方法改变，指向头结点，方便遍历
	var tail = new(Node)
	ls.headNode = tail //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i = i + 2 {
		var node = Node{data: i}
		//头插法,得到与插入顺序逆序的链表
		/* node.next = tail //将新插入的node的next指向头结点
		tail = node      //重新赋值头结点 */
		//ls.HeadInsert(&node)
		//尾插法
		/* 	tail.next = &node
		tail = &node */
		tail = tail.LastInsert(&node)
	}

	ls.Shownode() //遍历结果
	fmt.Println(ls.length())

	var ls02 = new(List) //不随插入方法改变，指向头结点，方便遍历
	var tail02 = new(Node)
	ls02.headNode = tail02
	for i := 2; i < 11; i = i + 2 {
		var node = Node{data: i}
		tail02 = tail02.LastInsert(&node)
	}
	ls02.Shownode() //遍历结果
	fmt.Println(ls02.length())

	mergeList(ls, ls02).Shownode()
}

//遍历
func (this *List) Shownode() {
	cur := this.headNode
	for cur != nil {
		fmt.Println(*cur)
		cur = cur.next
	}
}

//判断单链表是否为空
func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

//获取链表长度
func (this *List) length() int {
	cur := this.headNode
	count := 0

	for cur != nil {
		count++
		cur = cur.next
	}
	return count
}

//头插法
func (this *List) HeadInsert(p *Node) {
	p.next = this.headNode
	this.headNode = p
}

//尾插法
func (this *Node) LastInsert(p *Node) *Node {
	this.next = p
	this = p
	return this
}

//在指定位置插入元素
func (this *List) InsertByIndex(index int, p *Node) {

	pre := this.headNode
	count := 0
	for count < (index - 1) {
		pre = pre.next
		count++
	}

	p.next = pre.next
	pre.next = p

}

//删除指定位置的元素
func (this *List) RemoveAtIndex(index int) {
	pre := this.headNode
	count := 0
	for count != (index-1) && pre.next != nil {
		count++
		pre = pre.next
	}
	pre.next = pre.next.next
}

//删除指定值的元素
func (this *List) Remove(p *Node) {
	pre := this.headNode
	for pre.next != nil {
		if pre.next.data == p.data {
			pre.next = pre.next.next
		} else {
			pre = pre.next
		}
	}
}

//反转链表的实现
func reversrList(head *Node) *Node {
	cur := head
	var pre *Node = nil
	for cur != nil {
		//这句话最重要,go语言的并列赋值执行顺序为：
		//先计算=右边的变量值,再赋给左边对应的变量；
		//例如a,b=b,a，实则会交换顺序，注意千万不要拆开分步骤计算
		pre, cur, cur.next = cur, cur.next, pre
		//第一次循环,右边为：head,head.next,nil
		//左边为：pre,head,head.next,  即:pre=head,head=head.next,head.next=nil(三个运算并列执行，没有先后顺序，故前后相同变量值无联系)
	}
	return pre
}

//非递归合并两个有序链表
func mergeList(p1 *List, p2 *List) *List {
	var rt = new(List)
	var tmp = new(Node)
	rt.headNode = tmp
	p1node := p1.headNode.next
	p2node := p2.headNode.next
	for p1node != nil && p2node != nil {
		if p1node.data < p2node.data {
			tmp.next = p1node
			tmp = p1node
			p1node = p1node.next
		} else {
			tmp.next = p2node
			tmp = p2node
			p2node = p2node.next
		}
	}
	if p1node == nil {
		tmp.next = p2node
	} else {
		tmp.next = p1node
	}
	return rt
}

//链表排序
func sortList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	m := getMidean(head)
	right := sortList(m.next)
	m.next = nil
	left := sortList(head)
	return merge(left, right)
}
func merge(l1 *Node, l2 *Node) *Node {
	tmp := new(Node)
	tail := tmp
	for l1 != nil && l2 != nil {
		if l1.data < l2.data {
			tail.next = l1
			l1 = l1.next
		} else {
			tail.next = l2
			l2 = l2.next
		}
		tail = tail.next
	}
	if l1 != nil {
		tail.next = l1
	}
	if l1 != nil {
		tail.next = l2
	}
	return tmp.next

}
func getMidean(head *Node) *Node {
	slow := head
	fast := head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
