package lru_cache

type Node struct {
	key        int
	val        int
	next, prev *Node
}
type LRUCache struct {
	capacity   int
	hash       map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head := new(Node)
	tail := new(Node)
	head.val = -1
	head.key = -1
	tail.val = -1
	tail.key = -1

	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		hash:     make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) remove(key int) {
	node := this.hash[key]
	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = nil
	node.next = nil
	delete(this.hash, key)
}

func (this *LRUCache) insert(key, val int) {
	node := new(Node)
	node.val = val
	node.key = key

	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	node.prev = this.head

	this.hash[key] = node
}
func (this *LRUCache) Get(key int) int {
	if val, ok := this.hash[key]; ok {
		this.remove(key)
		this.insert(key, val.val)
		return val.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if this.Get(key) != -1 {
		//key found
		this.remove(key)
	}
	if this.capacity == len(this.hash) {
		this.remove(this.tail.prev.key)
	}
	this.insert(key, value)

}
