package main

import "fmt"

//想得到从小到大的排序结果，需要构建大顶堆，然后将堆顶最大值与最后的数据交换，
//依次进行，可以得到顺序的结果
func HeapSort(nums []int) {

	N := len(nums) - 1
	//注意，初始化堆时，需要从底部到顶部构建大（小）顶堆，最后一个非叶子节点开始
	var i int
	if N%2 == 0 {
		i = N/2 - 1
	} else {
		//最后一个非叶子节点只有左子树时
		i = N / 2
	}
	for i >= 0 {
		sink001(nums, i, N)
		i--
	}

	//将堆顶值和末尾交换，重新调整堆
	for j := N; j >= 0; j-- {
		wap(nums, 0, j)
		sink001(nums, 0, j-1) //交换之后，数组最后一位不算在堆内，需要减1操作
	}

}

//堆化,注意当初次堆化时,k为最后一个非叶子节点;当删除或者交换最大(小)堆顶调整时,k为0
func sink001(nums []int, k, N int) {
	for {
		i := 2*k + 1
		if i > N {
			break
		}
		//找左右子节点最大值
		if i+1 <= N && nums[i+1] < nums[i] {
			i++
		}
		//已经大于最大值，不需要再交换
		if nums[k] <= nums[i] {
			break
		}
		wap(nums, k, i)
		k = i //继续向下调整
	}
}

func wap(nums []int, x, y int) {
	nums[x], nums[y] = nums[y], nums[x]
}

func main() {
	array := []int{-1, -70, 0, 2, 1, 5, 8, 5, 6, -9, -11, 3, 2, 7, 10, 40, 60, 41, -2, -50}
	res := make([]int, 10)
	for i := 0; i < 10; i++ {
		res[i] = array[i]
	}

	N := 9
	//初次堆化，从底部到顶部构建大顶堆，最后一个非叶子节点开始
	var i int
	if N%2 == 0 {
		i = N/2 - 1
	} else {
		i = N / 2
	}
	for i >= 0 {
		sink001(res, i, N)
		i--
	}

	fmt.Print(res)

}
