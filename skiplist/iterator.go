package skiplist

type Iterator struct {
	list *SkipList
	node *Node
}

func (i *Iterator) Valid() bool {
	return i.node != nil
}

func (i *Iterator) Key() interface{} {
	return i.node.key
}

func (i *Iterator) Next() {
	i.list.mu.Lock()
	defer i.list.mu.Unlock()

	i.node = i.node.getNext(0)
}

func (i *Iterator) Prev() {
	i.list.mu.Lock()
	defer i.list.mu.Unlock()
	i.node = i.list.findLessThan(i.node.key)
	if i.node == i.list.head {
		i.node = nil
	}
}

func (i *Iterator) Seek(target interface{}) {
	i.list.mu.Lock()
	defer i.list.mu.Unlock()
	i.node, _ = i.list.findGreaterOrEqualKey(target)
}

// Position at the first entry in list.
// Final state of iterator is Valid() iff list is not empty.
func (i *Iterator) SeekToFirst() {
	i.list.mu.RLock()
	defer i.list.mu.RUnlock()

	i.node = i.list.head.getNext(0)
}

// Position at the last entry in list.
// Final state of iterator is Valid() iff list is not empty.
func (i *Iterator) SeekToLast() {
	i.list.mu.RLock()
	defer i.list.mu.RUnlock()

	i.node = i.list.findLast()
	if i.node == i.list.head {
		i.node = nil
	}
}
