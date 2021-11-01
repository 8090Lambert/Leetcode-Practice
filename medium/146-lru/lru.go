package _46_lru

type LRUCache struct {
	size int
	dict map[int]*Node
	head, tail *Node
}

type Node struct {
	key, val int
	prev, next *Node
}

func genNode(key, val int) *Node {
	n := new(Node)
	n.key = key
	n.val = val
	return n
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{}
	l.size = capacity
	l.dict = make(map[int]*Node)
	l.head = genNode(0, 0)
	l.tail = genNode(0, 0)
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.dict[key]; !ok {
		return -1
	}
	node := this.dict[key]
	this.moveToHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int)  {
	if v, ok := this.dict[key]; ok {
		v.val = value
		this.moveToHead(v)
	} else {
		node := genNode(key, value)
		this.dict[key] = node
		this.addHead(node)
		if len(this.dict) > this.size {
			remove := this.removeTail()
			delete(this.dict, remove.key)
		}
	}
}

func (this *LRUCache) removeTail() *Node{
	removed := this.tail.prev
	this.removeNode(removed)
	return removed
}

func (this *LRUCache) moveToHead(n *Node) {
	this.removeNode(n)
	this.addHead(n)
}

func (this *LRUCache) removeNode(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (this *LRUCache) addHead(n *Node) {
	this.head.next.prev = n
	n.next = this.head.next
	n.prev = this.head
	this.head.next = n
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */