package main

// 自己的代码，利用golang自带的slice来实现栈，具体是利用切片和append来模拟出栈和入栈，
type CQueue struct {
	// 用于存数字
	stack1 []int
	// 用于取数字
	stack2 []int
}
func Constructor_09() CQueue {
    return CQueue{
		stack1: []int{},
		stack2: []int{},
	}
}
func (this *CQueue) AppendTail(value int)  {
    this.stack1 = append(this.stack1, value)
}
func (this *CQueue) DeleteHead() int {
    if len(this.stack2) != 0{
    	value := this.stack2[len(this.stack2)-1]
    	this.stack2 = this.stack2[:len(this.stack2)-1]
    	return value
	} else {
		if len(this.stack1) != 0{
			for len(this.stack1) > 0 {
				value := this.stack1[len(this.stack1)-1]
				this.stack2 = append(this.stack2, value)
				this.stack1 = this.stack1[:len(this.stack1)-1]
			}
			value := this.stack2[len(this.stack2)-1]
			this.stack2 = this.stack2[:len(this.stack2)-1]
			return value
		} else {
			return -1
		}
	}
}

func main(){
    obj := Constructor_09()
    obj.AppendTail(5)
}


// 官方解法，利用了golang自带的双向链表list.List来实现，

//type CQueue struct {
//	stack1, stack2 *list.List
//}
//
//func Constructor() CQueue {
//	return CQueue{
//		stack1: list.New(),
//		stack2: list.New(),
//	}
//}
//
//func (this *CQueue) AppendTail(value int)  {
//	this.stack1.PushBack(value)
//}
//
//func (this *CQueue) DeleteHead() int {
//	// 如果第二个栈为空
//	if this.stack2.Len() == 0 {
//		for this.stack1.Len() > 0 {
//			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
//		}
//	}
//	if this.stack2.Len() != 0 {
//		e := this.stack2.Back()
//		this.stack2.Remove(e)
//		return e.Value.(int)
//	}
//	return -1
//}
//
//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/solution/mian-shi-ti-09-yong-liang-ge-zhan-shi-xian-dui-l-3/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
