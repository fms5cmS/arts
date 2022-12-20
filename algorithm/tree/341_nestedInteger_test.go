package tree

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

// 定义，这里只是为了本地代码不报错
type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool           { return false }
func (this NestedInteger) GetInteger() int           { return 0 }
func (n *NestedInteger) SetInteger(value int)        {}
func (this *NestedInteger) Add(elem NestedInteger)   {}
func (this NestedInteger) GetList() []*NestedInteger { return nil }

// NestedInteger 可以看作是一个 N 叉树，如 [[1,1],2,[1,1]] 可以看作：
//                       虚拟根节点
//          列表类型         2          列表类型
//       1          1               1          1
//
//  可以看出，这棵树的叶子节点对应一个整数，而非叶子节点则对应一个列表

type NestedIterator_Easy struct {
	vals []int
}

// 使用 DFS 将所有叶子节点保存到数组中，O(n)
func ConstructorForIterator_Easy(nestedList []*NestedInteger) *NestedIterator_Easy {
	vals := make([]int, 0)
	var dfs func(list []*NestedInteger)
	dfs = func(list []*NestedInteger) {
		for _, integer := range list {
			if integer.IsInteger() {
				vals = append(vals, integer.GetInteger())
			} else {
				dfs(integer.GetList())
			}
		}
	}
	dfs(nestedList)
	return &NestedIterator_Easy{vals: vals}
}

// O(1)
func (this *NestedIterator_Easy) Next_Easy() int {
	val := this.vals[0]
	this.vals = this.vals[1:]
	return val
}

// O(1)
func (this *NestedIterator_Easy) HasNext_Easy() bool {
	return len(this.vals) > 0
}

// 上面的解法虽然正确，但并不符合迭代器的定义。一般而言，迭代器是惰性的，也就是说，你需要一个结果，我就计算一个出来，而不是一次性把所有的结果都算出来

type NestedIterator struct {
	stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack := make([]*NestedInteger, 0, len(nestedList))
	// 将元素从后开始一个个入栈，因为栈是后进的先出！
	for i := len(nestedList) - 1; i >= 0; i-- {
		stack = append(stack, nestedList[i])
	}
	return &NestedIterator{stack: stack}
}

// 迭代器在调用 Next() 之前，都已经调用 HasNext() 进行过判断了
func (this *NestedIterator) Next() int {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return top.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	this.stackTop2Integer()
	return len(this.stack) > 0
}

// 将栈顶元素变成 integer
func (this *NestedIterator) stackTop2Integer() {
	for len(this.stack) > 0 {
		length := len(this.stack)
		top := this.stack[length-1]
		// 注意：只有在栈顶元素为 integer 时才会 return
		if top.IsInteger() {
			return
		}
		// 栈顶元素是列表，需要栈顶的列表先出栈，然后把列表中的元素从后往前加入栈中
		this.stack = this.stack[:length-1]
		list := top.GetList()
		for i := len(list) - 1; i >= 0; i-- {
			this.stack = append(this.stack, list[i])
		}
	}
}
