package skiplist

type Node struct {
	key  interface{}  // key,value存在一起
	next []*Node      // 列向指针
}

func NewNode(key interface{}, height int) *Node {
	n := new(Node)
	n.key = key
	n.next = make([]*Node, height)
	return n
}

func (n *Node) getNext(level int) *Node {
	return n.next[level]
}

func (n *Node) setNext(level int, node *Node) {
	n.next[level] = node
}
