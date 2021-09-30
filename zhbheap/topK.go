package main

type KthLargest struct {
	n    int
	topN []int
}

func Constructor(k int, nums []int) KthLargest {
	res := new(KthLargest)
	res.n = k
	res.topN = append(res.topN, nums...)
	if k > len(res.topN) {
		return *res
	}
	res.topN = res.topN[:k]
	var n int
	if (k-1)%2 == 0 {
		n = (k-1)/2 - 1
	} else {
		n = (k - 1) / 2
	}
	for n >= 0 {
		for {
			i := 2*n + 1
			if i > k-1 {
				break
			}
			if i+1 <= k-1 && res.topN[i+1] < res.topN[i] {
				i++
			}

			if res.topN[n] <= res.topN[i] {
				break
			}
			res.topN[n], res.topN[i] = res.topN[i], res.topN[n]
			n = i
		}
		n--
	}

	for j := k; j < len(nums); j++ {
		res.Add(nums[j])
	}
	return *res
}

func (this *KthLargest) Add(val int) int {
	if this.n > len(this.topN) {
		this.topN = append(this.topN, val)
	} else {
		if val > this.topN[0] {
			this.topN[0] = val
		} else {
			return this.topN[0]
		}
	}

	k := 0
	for {
		i := 2*k + 1
		if i > this.n-1 {
			break
		}
		if i+1 <= this.n-1 && this.topN[i+1] < this.topN[i] {
			i++
		}

		if this.topN[k] <= this.topN[i] {
			break
		}
		this.topN[k], this.topN[i] = this.topN[i], this.topN[k]
		k = i
	}

	return this.topN[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
