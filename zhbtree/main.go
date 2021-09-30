package main

import (
	"fmt"
	"math"
)

//定义二叉树的节点
type Node struct {
	value int
	left  *Node
	right *Node
}

//功能：打印节点的值
//参数：nil
//返回值：nil
func (node *Node) Print() {
	fmt.Printf("%d ", node.value)
}

//功能：设置节点的值
//参数：节点的值
//返回值：nil
func (node *Node) SetValue(value int) {
	node.value = value
}

//功能：创建节点
//参数：节点的值
//返回值：nil
func CreateNode(value int) *Node {
	return &Node{value, nil, nil}
}

//功能：查找节点，利用递归进行查找
//参数：根节点，查找的值
//返回值：查找值所在节点
func (node *Node) FindNode(n *Node, x int) *Node {
	if n == nil {
		return nil
	} else if n.value == x {
		return n
	} else {
		p := node.FindNode(n.left, x)
		if p != nil {
			return p
		}
		return node.FindNode(n.right, x)
	}
}

//功能：求树的高度
//参数：根节点
//返回值：树的高度，树的高度=Max(左子树高度，右子树高度)+1
func (node *Node) GetTreeHeigh(n *Node) int {
	if n == nil {
		return 0
	} else {
		lHeigh := node.GetTreeHeigh(n.left)
		rHeigh := node.GetTreeHeigh(n.right)
		if lHeigh > rHeigh {
			return lHeigh + 1
		} else {
			return rHeigh + 1
		}
	}
}

//功能：递归前序遍历二叉树
//参数：根节点
//返回值：nil
func (node *Node) PreOrder(n *Node) {
	if n != nil {
		fmt.Printf("%d ", n.value)
		node.PreOrder(n.left)
		node.PreOrder(n.right)
	}
}

//功能：递归中序遍历二叉树
//参数：根节点
//返回值：nil
func (node *Node) InOrder(n *Node) {
	if n != nil {
		node.InOrder(n.left)
		fmt.Printf("%d ", n.value)
		node.InOrder(n.right)
	}
}

//功能：递归后序遍历二叉树
//参数：根节点
//返回值：nil
func (node *Node) PostOrder(n *Node) {
	if n != nil {
		node.PostOrder(n.left)
		node.PostOrder(n.right)
		fmt.Printf("%d ", n.value)
	}
}

//功能：打印所有的叶子节点
//参数：root
//返回值：nil
func (node *Node) GetLeafNode(n *Node) {
	if n != nil {
		if n.left == nil && n.right == nil {
			fmt.Printf("%d ", n.value)
		}
		node.GetLeafNode(n.left)
		node.GetLeafNode(n.right)
	}
}

func main() {
	//创建一颗树
	root := CreateNode(5)
	root.left = CreateNode(2)
	root.right = CreateNode(4)
	root.left.right = CreateNode(7)
	root.left.right.left = CreateNode(6)
	root.right.left = CreateNode(8)
	root.right.right = CreateNode(9)

	fmt.Printf("%d\n", root.FindNode(root, 4).value)
	fmt.Printf("%d\n", root.GetTreeHeigh(root))

	root.PreOrder(root)
	fmt.Printf("\n")
	root.InOrder(root)
	fmt.Printf("\n")
	root.PostOrder(root)
	fmt.Printf("\n")

	root.GetLeafNode(root)
	fmt.Printf("\n")

	i := make([]int, 0)
	i = append(i, 1)
	i = append(i, 2)
	i = append(i, 3)
	i = i[3:]
	fmt.Print(i)

}

//二叉树节点数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := root.Left
	right := root.Right
	lh := 0
	rh := 0
	for left != nil {
		lh++
		left = left.Left
	}
	for right != nil {
		rh++
		right = right.Right
	}

	if lh == rh {
		return int(math.Pow(2, float64(lh)+1)) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

//翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
