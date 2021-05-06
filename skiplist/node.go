package skiplist

type Node struct {
	key  interface{}
	next []*Node
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
