package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//func getLength(head *ListNode) (length int) {
//	for ; head != nil; head = head.Next {
//		length++
//	}
//	return
//}
//
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	length := getLength(head)
//	// 哑节点
//	dummy := &ListNode{0, head}
//	cur := dummy
//	for i := 0; i < length-n; i++ {
//		cur = cur.Next
//	}
//	cur.Next = cur.Next.Next
//	return dummy.Next
//}

//func getLength(head *ListNode)  (i int){
//	for ; head != nil; i++{
//		head = head.Next
//	}
//	return
//}
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	length := getLength(head)
//	dummy := new(ListNode)
//	dummy.Next = head
//	transcript := dummy
//	for i := 0; i < length-n; i++{
//		transcript = transcript.Next
//	}
//	transcript.Next = transcript.Next.Next
//	return dummy.Next
//}

//利用栈求解

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	all_node := [] *ListNode{}
	dummpy := &ListNode{0, head}
	for k := dummpy; k != nil; k = k.Next{
		all_node = append(all_node, k)
	}
	del_pre := all_node[len(all_node)-1-n]
	del_pre.Next = del_pre.Next.Next
	return dummpy.Next
}
//https://www.cnblogs.com/nima/p/12724826.html  链表的生成
func main()  {
	node1 := new(ListNode)
	node1.Val = 20
	node2 := new(ListNode)
	node2.Val = 50
	node1.Next = node2
	node3 := new(ListNode)
	node3.Val = 30
	node2.Next = node3
	res := removeNthFromEnd(node1, 2)
	for {
		if res != nil{
			println(res.Val)
			res = res.Next
		}else {
			break
		}
	}
}
