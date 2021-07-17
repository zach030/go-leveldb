package memtable

type ValueType int

const (
	TypeDeletion ValueType = 0
	TypeValue ValueType = 1
)

type InternalKey struct {
	b []byte
}

func newInternalKey(seq int64,valueType ValueType,key,value []byte)*InternalKey{
	return &InternalKey{}
}