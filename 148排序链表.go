package main

//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func merge(h1 *ListNode, h2 *ListNode) *ListNode {
	// 设置一个哑节点，
	res_copy := &ListNode{Val: 0}
	res := res_copy
	for ; h1 != nil && h2 != nil; {
		if h1.Val < h2.Val{
			res.Next = h1
			h1 = h1.Next
			// 刚开始丢了这句是错的，因为res是链表尾巴，也要跟着走，
			res = res.Next
		} else {
			res.Next = h2
			h2 = h2.Next
			res = res.Next
		}
	}
	if h1 != nil{
		res.Next = h1
	} else {
		res.Next = h2
	}
	return res_copy.Next
}
func mid(h *ListNode) (*ListNode, bool) {
	//if h == nil{
	//	return nil, true
	//}
	memo := h
	res := h
	// 如果for循环结束后想要用结束前一轮的值，就需要定义全局变量来保存需要的值，
	k := &ListNode{}
	for ; res != nil && res.Next != nil; {
		k = h
		h = h.Next
		res = res.Next.Next
	}
	// 找到链表中点后，务必要断开这个链表，否则下次调用mid函数的时候，就会超出去，
    k.Next = nil
	if memo == h{
		return h, true
	}else {
		return h, false
	}
}
func sortList(head *ListNode) *ListNode {
	mid_postion, sign := mid(head)
	// 如果不能再分割了，就直接返回节点，否则继续分，而一旦继续再分，必然是一个重复的过程，就用到了递归，
	// 左右两边各自递归完，必然返回的都是各自的头，之后再进行归并，
	if sign{
		return head
	} else {
		h := sortList(head)
		m := sortList(mid_postion)
		res := merge(h, m)
		return res
	}
}

func main()  {
	h := ListNode{Val: 4}
	h.Next = &ListNode{Val: 2}
	h.Next.Next = &ListNode{Val: 1}
	h.Next.Next.Next = &ListNode{Val: 3}
	res := sortList(&h)
	println(res)
}