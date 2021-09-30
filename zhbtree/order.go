package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历，入栈时访问数据
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	//树的节点记录栈
	stack := make([]*TreeNode, 0)
	//树的节点遍历数据存放数组
	res := make([]int, 0)
	p := root
	//退出循环条件为：当节点和栈为空同时为空时
	for p != nil || len(stack) != 0 {
		//如果节点不为空,该节点进栈，并访问该节点数据存入数组，向左节点继续重复过程，直到节点为空
		if p != nil {
			stack = append(stack, p)
			res = append(res, p.Val)
			p = p.Left
		} else {
			//当节点为空时,节点回退到栈顶, 然后向右子树继续；此时栈顶还要出栈, 避免len(stack) == 0循环退出条件一直不成立
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	return res
}

//中序遍历,出栈时访问数据
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	p := root
	for p != nil || len(stack) != 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			res = append(res, p.Val)
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	return res
}

//后序遍历，出栈时访问数据
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	//新增一个对栈元素的访问次数计数的栈,第一次访问该元素时置位1,当栈顶元素的flagstack为1时，说明是从左子树回退，则把flagstack置位2，先右子树继续。当栈顶元素的flagstack为2时，则说明是从右子树来的，则进行出栈和访问该元素数据。
	flagStack := make([]int, 0)
	res := make([]int, 0)
	p := root
	for p != nil || len(stack) != 0 {
		if p != nil {
			stack = append(stack, p)
			flagStack = append(flagStack, 1)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			if flagStack[len(flagStack)-1] == 1 {
				flagStack[len(flagStack)-1] = 2
				p = p.Right
			} else {
				res = append(res, p.Val)
				stack = stack[:len(stack)-1]
				flagStack = flagStack[:len(flagStack)-1]
				p = nil
			}
		}
	}
	return res
}

//已知前序和中序构造二叉树，返回根节点
func buildTree(preorder []int, inorder []int) *TreeNode {
	//如果遍历数组为空,则返回节点为nil
	if len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	//如果遍历数组长度为1,则只有一个节点,直接返回
	if len(preorder) == 1 {
		return root
	}
	//如果遍历数组长度在2以上时,进入递归：先根据前序的第一个节点（根节点）把中序数组分成两部分左子树和右子树;再分别对左右子树的前序中序遍历继续求孙子树...
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			if i == len(inorder)-1 {
				//当i为中序最后一个节点时,则二叉树只有左子树,故右子树需置位nil
				root.Left = buildTree(preorder[1:], inorder[0:i])
				root.Right = buildTree(nil, nil)
			} else {
				//包含一种特例，当i为0时,二叉树只有右子树
				root.Left = buildTree(preorder[1:i+1], inorder[0:i])
				root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			}
			break
		}
	}
	return root
}

//中序和后续构造二叉树,实现与上面完全一致
func buildTree2(inorder []int, postorder []int) *TreeNode {
	lp := len(postorder)
	if lp == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = postorder[lp-1]
	if lp == 1 {
		return root
	}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[lp-1] {
			if i == len(inorder)-1 {
				root.Left = buildTree2(inorder[0:i], postorder[0:i])
				root.Right = buildTree2(nil, nil)
			} else {
				root.Left = buildTree2(inorder[0:i], postorder[0:i])
				root.Right = buildTree2(inorder[i+1:], postorder[i:lp-1])
			}
		}
	}
	return root
}

/* 层序遍历,队列实现：
（1）根结点入队；
（2）当队列不空时，重复下列操作：
   从队列退出一个结点并访问；
若其有左孩子，则访问左孩子，并将其左孩子入队；
若其有右孩子，则访问右孩子，并将其右孩子入队； */
func levelOrder(root *TreeNode) [][]int {
	//构造节点指针数组,则用数组实现队列
	q := make([]*TreeNode, 0)
	//构造返回值的二维数组
	res := make([][]int, 0)

	if root == nil {
		return res
	}
	//初始根节点入队列
	q = append(q, root)
	//当队列不为空时
	for len(q) != 0 {
		//进行该层节点数统计,则一维数组长度
		size := len(q)
		//构建一维数组
		tmp := make([]int, 0, size)
		for i := 0; i < size; i++ {
			front := q[0]
			//头节点出队列操作
			q = q[1:]
			tmp = append(tmp, front.Val)
			if front.Left != nil {
				q = append(q, front.Left)
			}
			if front.Right != nil {
				q = append(q, front.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}
