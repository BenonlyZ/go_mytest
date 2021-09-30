package LinkedNodes

import "fmt"

//头节点,为了遍历使用
var heads *Nodes

//当前节点
var currs *Nodes

//节点结构体里的数据信息,存放该节点相应的key和value
type MP struct {
	K string
	V string
}

//节点结构体
type Nodes struct {
	Data     MP
	NextNode *Nodes
}

//创建哈希表链表部分每个链表的头节点,则为数组部分的数组元素
func CreateArryNode(k, v string) *Nodes {
	//创建Node头节点
	var node = new(Nodes)

	node.Data.V = v
	node.Data.K = k
	node.NextNode = nil

	//头结点和当前节点均初始为刚创建的节点
	heads = node
	currs = node

	return node
}

//向指定的节点添加新节点
func AddChilddNode(k, v string, currs *Nodes) *Nodes {
	var newNode *Nodes = new(Nodes)
	//添加信息
	newNode.Data.K = k
	newNode.Data.V = v
	newNode.NextNode = nil
	//挂接节点
	currs.NextNode = newNode
	currs = newNode
	//fmt.Println(curr)
	return newNode
}

//遍历指定的节点链表
func ShowNode(n *Nodes) {
	var node = n
	for node != nil {
		fmt.Println(node.Data)
		node = node.NextNode
	}
}

//计算节点个数
func NodesCount() int {
	var n = heads
	var flag int
	for n != nil {
		flag += 1
		n = n.NextNode
	}
	return flag
}
