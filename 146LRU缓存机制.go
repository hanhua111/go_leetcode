package main

// 自己的代码
type Node struct {
	// 添加key是为了在Put方法中删除tail的前一个节点时，同时删除dict中key对应的node，
	key    int
	value  int
	next   *Node
	prev   *Node
}
type LRUCache struct {
    cap           int
    count         int
    dict          map[int]*Node
    linked_list   *LinkedList
}
type LinkedList struct {
	head *Node
	tail *Node
}
func NewLinkedList()  *LinkedList {
	//var tail *Node
	//head := &Node{
	//	value: 0,
	//	// 注意这里的tail是nil，后面还需要再赋值一次，
	//	next:  tail,
	//	prev:  nil,}
	//tail = &Node{
	//	value: 0,
	//	next:  nil,
	//	prev:  head,
	//}
	//res := &LinkedList{
	//	head: head,
	//	tail: tail,
	//}
	//res.head.next = tail

	// 这样写更简洁，先同时初始化后，再连接二者，
	res := &LinkedList{
		head: &Node{value:0},
		tail: &Node{value:0},
	}
	res.head.next = res.tail
	res.tail.prev = res.head
	return res
}
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:          capacity,
        dict:         make(map[int]*Node),
        linked_list:  NewLinkedList(),
        count:        0,
	}
}
// 不用用key，用node地址更好，
func (this *LRUCache) Delete(n *Node)  {
	// 先删除该node节点
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (this *LRUCache) MoveToHead(key int){
	// 将该node节点移动到head节点后面，
	this.dict[key].next = this.linked_list.head.next
	this.dict[key].prev = this.linked_list.head
    //fmt.Println(this.linked_list.head.next)
	this.dict[key].next.prev = this.dict[key]
	this.linked_list.head.next = this.dict[key]
}

func (this *LRUCache) Get(key int) int {
    // 先检查有没有该值，如果有，移动到链表头，否则检查容量，再插入，
    if _, ok := this.dict[key]; ok{
    	this.Delete(this.dict[key])
    	this.MoveToHead(key)
	} else {
		return -1
	}
	return this.dict[key].value
}
func (this *LRUCache)Judge() bool {
	if this.count < this.cap{
		return true
	} else {
		return false
	}
}
func (this *LRUCache) Put(key int, value int)  {
    // 先检查dict中有没有那个key，有直接移动，没有的话先检查容量，再插入，
    if _, ok := this.dict[key]; ok{
    	// 移动之前要先删除该节点，再移动，
    	this.Delete(this.dict[key])
    	this.MoveToHead(key)
    	this.dict[key].value = value
	} else {
		insert := &Node{
			key:   key,
			value: value,
			//next:  this.linked_list.head.next,
			next:  nil,
			//prev:  this.linked_list.head,
			prev:  nil,
		}
		this.dict[key] = insert
		if this.Judge(){
			// 直接插入
			this.MoveToHead(key)
		} else {
			// 否则先删除dict中的key，再删除linked_list的前一个，再插入head后面，
			// 删除dict中key和双向链表中node的顺序不可颠倒，因为要利用链表中倒数第二个node去寻找dict的key，
			delete(this.dict, this.linked_list.tail.prev.key)
			this.Delete(this.linked_list.tail.prev)
			this.MoveToHead(key)
		}
		this.count += 1
	}
}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// 思路：用一个字典存key和链表节点的地址，用一个双向链表用于存储实际的value值，
// 在双向链表中，将一个node节点移动到head节点后面要做六次修改，首先删除该node节点做两次修改，再插入head后要做六次修改，
// 这个题的关键之处就在于删除node和移动node，最好是把这两个操作分开写成两个函数，实现解耦，
func main()  {
	capacity := 2
	obj := Constructor(capacity)
	obj.Put(2,1)
	obj.Put(1,1)
	obj.Put(2,3)
	obj.Put(4,1)
	obj.Get(1)
	obj.Get(2)
	//obj.Put(1, 1) // 缓存是 {1=1}
	//obj.Put(2, 2) // 缓存是 {1=1, 2=2}
	//obj.Get(1)           // 返回 1
	//obj.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	//obj.Get(2)           // 返回 -1 (未找到)
	//obj.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	//obj.Get(1)           // 返回 -1 (未找到)
	//obj.Get(3)           // 返回 3
	//obj.Get(4)           // 返回 4
}



//官方解法

//type LRUCache struct {
//	size int
//	capacity int
//	cache map[int]*DLinkedNode
//	head, tail *DLinkedNode
//}
//
//type DLinkedNode struct {
//	key, value int
//	prev, next *DLinkedNode
//}
//
//func initDLinkedNode(key, value int) *DLinkedNode {
//	return &DLinkedNode{
//		key: key,
//		value: value,
//	}
//}
//
//func Constructor(capacity int) LRUCache {
//	l := LRUCache{
//		cache: map[int]*DLinkedNode{},
//		head: initDLinkedNode(0, 0),
//		tail: initDLinkedNode(0, 0),
//		capacity: capacity,
//	}
//	l.head.next = l.tail
//	l.tail.prev = l.head
//	return l
//}
//
//func (this *LRUCache) Get(key int) int {
//	if _, ok := this.cache[key]; !ok {
//		return -1
//	}
//	node := this.cache[key]
//	this.moveToHead(node)
//	return node.value
//}
//
//
//func (this *LRUCache) Put(key int, value int)  {
//	if _, ok := this.cache[key]; !ok {
//		node := initDLinkedNode(key, value)
//		this.cache[key] = node
//		this.addToHead(node)
//		this.size++
//		if this.size > this.capacity {
//			removed := this.removeTail()
//			delete(this.cache, removed.key)
//			this.size--
//		}
//	} else {
//		node := this.cache[key]
//		node.value = value
//		this.moveToHead(node)
//	}
//}
//
//func (this *LRUCache) addToHead(node *DLinkedNode) {
//	node.prev = this.head
//	node.next = this.head.next
//	this.head.next.prev = node
//	this.head.next = node
//}
//
//func (this *LRUCache) removeNode(node *DLinkedNode) {
//	node.prev.next = node.next
//	node.next.prev = node.prev
//}
//
//func (this *LRUCache) moveToHead(node *DLinkedNode) {
//	this.removeNode(node)
//	this.addToHead(node)
//}
//
//func (this *LRUCache) removeTail() *DLinkedNode {
//	node := this.tail.prev
//	this.removeNode(node)
//	return node
//}

















































