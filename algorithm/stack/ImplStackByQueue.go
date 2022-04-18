package stack

/**
请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。

实现 MyStack 类：

void push(int x) 将元素 x 压入栈顶。
int pop() 移除并返回栈顶元素。
int top() 返回栈顶元素。
boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。


注意：
你只能使用队列的基本操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。

自我总结：
1、为什么需要两个队列，因为要考虑到题目要求，只能遵循队列的先进先出特性

2、那我们可以不可以用一个队列来实现栈呢？答案是肯定的。
我们只需要执行以下两个步骤就可以实现将队列转换为栈了，具体实现步骤如下：
将元素入列到队尾；
再将除队尾之外的所有元素移除并重写入列。
这样操作之后，最后进入的队尾元素反而变成了队头元素，也就实现了后进先出的功能了，
*/

type MyStack struct {
	enqueue []int
	dequeue []int
}

func Constructor() MyStack {
	return MyStack{enqueue: []int{}, dequeue: []int{}}
}

func (this *MyStack) Push(x int) {
	this.enqueue = append(this.enqueue, x)
}

func (this *MyStack) Pop() int {
	length := len(this.enqueue)
	for i:=0;i<length-1;i++{
		this.dequeue = append(this.dequeue, this.enqueue[0])
		this.enqueue = this.enqueue[1:]
	}

	topElement := this.enqueue[0]
	this.enqueue = this.dequeue
	this.dequeue = nil

	return topElement
}

func (this *MyStack) Top() int {
	topElement := this.Pop()
	this.enqueue = append(this.enqueue, topElement)
	return topElement
}

func (this *MyStack) Empty() bool {
	if len(this.enqueue) == 0{
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
