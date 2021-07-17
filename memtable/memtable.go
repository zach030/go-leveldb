package memtable

import (
	"go-leveldb/skiplist"
	"go-leveldb/utils"
)

type MemTable struct {
	table *skiplist.SkipList
}

func NewMemTable() *MemTable {
	return &MemTable{
		table: skiplist.New(utils.IntComparator),
	}
}

func (m *MemTable) NewIterator() *skiplist.Iterator {
	return m.table.NewIterator()
}

func (m MemTable) Add(seq int64, valueType ValueType, key, value []byte) {
	// encode
	m.table.Insert(newInternalKey(seq, valueType, key, value))
}
