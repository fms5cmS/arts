package stackAndQueueRela

// 155. 最小栈：设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
type MinStack struct {
	// 栈顶指针
	top int
	// 存放元素、存放当前栈中的最小值
	data, min []int
}

func ConstructorMinStack() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
		top:  -1,
	}
}

// 第一次压栈时，同时向 data 和 min 中压入第一个元素；
// 其他时候，data 正常压栈，min 有所区别：
//    比较 min 中栈顶元素与当前压入 data 中元素 X 的大小，
//		1. 如果 X 较小，则向 min 中也压入 X
//		2. 如果 min 的栈顶元素较小，则重复压入 min 的栈顶元素
func (this *MinStack) Push(x int) {
	this.top++
	this.data = append(this.data, x)
	if this.top != 0 && x > this.min[this.top-1] {
		this.min = append(this.min, this.min[this.top-1])
		return
	}
	this.min = append(this.min, x)
}

// 每次 data 进行 pop 操作时，min 也要进行 pop 操作，
// 这样 min 的栈顶元素就始终是当前 data 中所有元素的最小值
func (this *MinStack) Pop() {
	this.data = this.data[:this.top]
	this.min = this.min[:this.top]
	this.top--
}

// 获取 data 的栈顶元素
func (this *MinStack) Top() int {
	return this.data[this.top]
}

// 获取 min 的栈顶元素即可
func (this *MinStack) GetMin() int {
	return this.min[this.top]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
