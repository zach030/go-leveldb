package memtable

import (
	"errors"
	"go-leveldb/skiplist"
	"go-leveldb/utils"
)

type MemTable struct {
	table *skiplist.SkipList
}

func NewMemTable() *MemTable {
	return &MemTable{
		table: skiplist.New(InternalKeyComparator),
	}
}

func (m *MemTable) NewIterator() *skiplist.Iterator {
	return m.table.NewIterator()
}

func (m MemTable) Add(seq int64, valueType ValueType, key, value []byte) {
	// encode
	m.table.Insert(newInternalKey(seq, valueType, key, value))
}

func (m *MemTable) Get(key []byte) (bool, []byte, error) {
	// 构造存储的key entry
	lookUpKey := LookUpKey(key)
	it := m.NewIterator()
	it.Seek(lookUpKey)
	if it.Valid() {
		internalKey := it.Key().(*InternalKey)
		if utils.UserKeyComparator(key, internalKey.userKey()) == 0 {
			// value type
			if internalKey.valueType() == TypeValue {
				return true, internalKey.userValue(), nil
			}
			return true, nil, errors.New("not found")
		}
	}
	return true, nil, errors.New("not found")
}

func InternalKeyComparator(a, b interface{}) int {
	// 先按userKey排序，再按照seq排序
	aK := a.(*InternalKey)
	bK := b.(*InternalKey)
	ret := utils.UserKeyComparator(aK.b, bK.b)
	if ret == 0 {
		anum := aK.seq()
		bnum := bK.seq()
		if anum > bnum {
			ret = -1
		} else if anum < bnum {
			ret = 1
		}
	}
	return ret
}
