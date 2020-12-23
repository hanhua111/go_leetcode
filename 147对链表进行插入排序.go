package main

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

//// 链表不适合二分查找，因为它和数组的特性不一样，与其二分查找，不如直接暴力找，
//// 自己的二分查找代码，写的有问题，和用暴力遍历时间复杂度一样，用暴力更简洁，
//var Head = &ListNode{}
//var Tail = &ListNode{}
//// 利用快慢指针二分来找到要插入的位置，
//func mid_(h *ListNode, t *ListNode) *ListNode {
//	//res := &ListNode{}
//	k := h
//	for ; k.Next != t && k.Next.Next != t; {
//		k = k.Next.Next
//		h = h.Next
//	}
//	return h
//}
//// 找到位置后进行插入，
//func insert_ele(ele *ListNode) {
//	if Head == nil{
//		Head = ele
//		Tail = ele
//	} else {
//		for ; Head != Tail;{
//			Mid := mid_(Head, Tail)
//			if Mid.Val > ele.Val{
//				Tail = Mid
//			} else {
//				Head = Mid
//			}
//		}
//	}
//}
//func insertionSortList(head *ListNode) *ListNode {
//	for ; head != nil; {
//		insert_ele(head)
//		head = head.Next
//	}
//	return Head
//}

func main()  {

}
























