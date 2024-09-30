package main

func main() {

}

type CustomStack struct {
	arr     []int
	index   int
	maxSize int
}

func Constructor(maxSize int) CustomStack {
	c := CustomStack{}
	c.arr = make([]int, maxSize)
	c.index = 0
	c.maxSize = maxSize

	return c
}

func (this *CustomStack) Push(x int) {
	if this.index == this.maxSize {
		return
	}

	this.arr[this.index] = x
	this.index += 1
}

func (this *CustomStack) Pop() int {
	if this.index == 0 {
		return -1
	}

	res := this.arr[this.index-1]
	this.arr[this.index-1] = -1
	this.index -= 1

	return res
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k; i++ {
		if i >= this.index {
			break
		}

		this.arr[i] += val
	}
}
