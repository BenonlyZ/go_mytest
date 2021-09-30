package main

import (
	"fmt"
	"math"
	"myhash/HashMap"
)

type MyHashMap struct {
	Node []struct {
		k int
		v int
	}
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {

	var myhash = MyHashMap{
		Node: make([]struct {
			k int
			v int
		}, 16),
	}
	//初始化时，当key为0,且myhash.Node[0].k=0说明已插入key 0,与原意不符合,故k置位非0
	myhash.Node[0].k = -1
	myhash.Node[0].v = -1
	return myhash
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	size := len(this.Node)
	if key >= size {
		max := math.Max(float64(key+1), float64(size*2))
		newhash := make([]struct {
			k int
			v int
		}, int(max))
		//因为newhash的初始长度为max,使用append会在new[hash-1]后追加，不能达到按下标拷贝的效果
		for i := 0; i < size; i++ {
			newhash[i] = this.Node[i]
		}
		this.Node = newhash
	}
	this.Node[key].k = key
	this.Node[key].v = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if key < len(this.Node) && key == this.Node[key].k {
		return this.Node[key].v
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	if key < len(this.Node) && key == this.Node[key].k {
		this.Node[key].k = -1
		this.Node[key].v = -1
	}
}

func main() {

	//创建数组，添加数组元素，则链表的头节点
	HashMap.CreateArry()

	//随机向数组的每个下标添加子节点
	HashMap.AddKVToArr("abc", "世界")
	HashMap.AddKVToArr("def", "和平")

	HashMap.GetValue("abc1")
	hash := Constructor()
	hash.Put(2, 34)
	hash.Put(3, 78)
	fmt.Println(hash.Get(2))
	hash.Remove(2)
	fmt.Println(hash.Get(2))
}
