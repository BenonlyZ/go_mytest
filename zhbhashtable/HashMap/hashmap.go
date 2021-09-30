package HashMap

import (
	"fmt"
	"myhash/LinkedNodes"
)

//定义哈希表数组部分的全局变量,每个数组元素为Nodes类型指针
var arr [16]*LinkedNodes.Nodes

//创建16个顶层节点，放到数组中
func CreateArry() {
	var ar = [16]*LinkedNodes.Nodes{}
	//这里头结点不存具体的key和value值
	for i := 0; i < 16; i++ {
		ar[i] = LinkedNodes.CreateArryNode("头节点", "头节点")
	}
	//赋值给全局变量
	arr = ar
	//fmt.Println(ar)
}

//向数组中添加键值对
func AddKVToArr(k, v string) {
	//先计算出要添加的数据存储到哪个下角标中，这里调用从网上找的算法
	var corner = HashCode(k)
	var head = arr[corner]
	//调用添加方法
	LinkedNodes.AddChilddNode(k, v, head)
}

//获取数据
func GetValue(k string) {
	//先判断是哪个下标存储
	var corner = HashCode(k)
	//获取头节点
	var head = arr[corner]
	//通过头节点遍历
	for head != nil {
		if head.Data.K == k {
			fmt.Println(head.Data.V)
			break
		} else {
			head = head.NextNode
		}
	}
	fmt.Println("该键不存在")
}

//将key转换成数组下标的散列算法，范围16之间
func HashCode(key string) int {
	var index int
	index = int(key[0])
	for k := 0; k < len(key); k++ {
		index *= (1103515245 + int(key[k]))
	}
	index >>= 27
	index &= 16 - 1

	return index
}
