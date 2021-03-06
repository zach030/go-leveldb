package skiplist

import (
	"go-leveldb/utils"
	"math/rand"
	"sync"
)

const (
	kMaxHeight = 12
	kBranching = 4
)

type SkipList struct {
	maxHeight  int
	head       *Node
	comparator utils.Comparator
	mu         sync.RWMutex
}

func New(comparator utils.Comparator) *SkipList {
	return &SkipList{
		maxHeight:  1,
		head:       NewNode(nil, kMaxHeight),
		comparator: comparator,
	}
}

func (s *SkipList) NewIterator() *Iterator {
	return &Iterator{
		list: s,
	}
}

func (s *SkipList) Insert(key interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, prev := s.findGreaterOrEqualKey(key)
	// 找到新插入后元素的每一层的前驱节点
	height := s.randomHeight()
	// 调整level
	if height > s.maxHeight {
		for i := s.maxHeight; i < height; i++ {
			// 新增的节点的前驱是head指针
			prev[i] = s.head
		}
		s.maxHeight = height
	}
	// 新建节点并插入
	h := NewNode(key, height)
	for i := 0; i < height; i++ {
		// 逐层插入新节点
		h.setNext(i, prev[i].getNext(i))
		prev[i].setNext(i, h)
	}
}

// 找到对应key值适当节点和遍历队列
func (s *SkipList) findGreaterOrEqualKey(key interface{}) (*Node, [kMaxHeight]*Node) {
	var prev [kMaxHeight]*Node
	h := s.head
	level := s.maxHeight - 1
	for true {
		// 同层横向查找,该层中下一个节点
		next := h.getNext(level)
		if s.keyIsAfterNode(key, next) {
			// 待查找 key 比 next 大，则在该层继续查找
			h = next
		} else {
			// 跳到下一级，记录prev指针（查找的路径）
			prev[level] = h
			// 找到最底层，next为目标节点
			if level == 0 {
				return next, prev
			}
			level--
		}
	}
	return nil, prev
}

// node节点的key比给定key小,则应该放在node后面
func (s *SkipList) keyIsAfterNode(key interface{}, n *Node) bool {
	return (n != nil) && (s.comparator(n.key, key) < 0)
}

func (s *SkipList) randomHeight() int {
	height := 1
	for height < kMaxHeight && (rand.Intn(kBranching) == 0) {
		height++
	}
	return height
}

func (s *SkipList) findLessThan(key interface{}) *Node {
	x := s.head
	level := s.maxHeight - 1
	for true {
		next := x.getNext(level)
		if next == nil || s.comparator(next.key, key) >= 0 {
			if level == 0 {
				return x
			} else {
				level--
			}
		} else {
			x = next
		}
	}
	return nil
}

func (s *SkipList) findLast() *Node {
	x := s.head
	level := s.maxHeight - 1
	for true {
		next := x.getNext(level)
		if next == nil {
			if level == 0 {
				return x
			} else {
				level--
			}
		} else {
			x = next
		}
	}
	return nil
}