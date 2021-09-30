package main

import "fmt"

type QueueNode struct{
	Data interface{}
	Next *QueueNode
}
//创建队列
func (q *QueueNode)Create(data ...interface{}){
	if q==nil{
		return
	}
	if len(data)==0{
		return
	}

	for _,v:=range data{
		newNode:=new(QueueNode)
		newNode.Data=v
		q.Next=newNode
		q=q.Next
	}
}
//打印队列
func (q *QueueNode)Print(){
	if q==nil{
		return 
	}
	for q!=nil{
		if q.Data!=nil{
			fmt.Print(q.Data," ")
		}
		q=q.Next
	}
}
//统计队列节点个数
func (q *QueueNode) Length()int{
  if q==nil{
	  return -1
  }
  i:=0
  for q.Next!=nil{
	  i++
	  q=q.Next
  }
  return i+1
}

//入列
func (q *QueueNode)Push(data interface{}){
	if q==nil{
		return
	}
	if data==nil{
		return
	}

	for q.Next!=nil{
		q=q.Next
	}

	newNode:=new(QueueNode)
	newNode.Data=data

	q.Next=newNode
}

//出队列
func (q *QueueNode)Pop(){
	if q==nil{
		return
	}
	q.Next=q.Next.Next
}

func main(){
	fmt.Println()
}