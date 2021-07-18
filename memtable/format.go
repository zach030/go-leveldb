package memtable

import "encoding/binary"

type ValueType int

const (
	TypeDeletion ValueType = 0
	TypeValue    ValueType = 1
)

type InternalKey struct {
	b []byte
}

func newInternalKey(seq int64, valueType ValueType, key, value []byte) *InternalKey {
	// format of entry
	internalKeySize := len(key) + 4
	valueSize := len(value)
	encodeLen := 4 + internalKeySize + 4 + valueSize
	buf := make([]byte, encodeLen)
	offset := 0
	binary.LittleEndian.PutUint32(buf[offset:], uint32(internalKeySize))
	offset += 4
	copy(buf[offset:], key)
	offset += len(key)
	binary.LittleEndian.PutUint64(buf[offset:], uint64(seq)<<8|uint64(valueType))
	offset += 8
	binary.LittleEndian.PutUint32(buf[offset:], uint32(valueSize))
	offset += 4
	copy(buf[offset:], value)
	return &InternalKey{b: buf}
}

func LookUpKey(key []byte) *InternalKey {
	buf := make([]byte, 4+len(key)+8)
	offset := 0
	binary.LittleEndian.PutUint32(buf[offset:], uint32(len(key)+8))
	offset += 4
	copy(buf[offset:], key)
	offset += len(key)
	binary.LittleEndian.PutUint64(buf[offset:], 0XFFFFFFFFFFFFFFFF)
	return &InternalKey{b: buf}
}

func (i *InternalKey) userKey() []byte {
	size := binary.LittleEndian.Uint32(i.b)
	return i.b[4 : size-4]
}

func (i *InternalKey) userValue() []byte {
	valueSize := binary.LittleEndian.Uint32(i.b) + 8
	return i.b[valueSize:]
}

func (i *InternalKey) valueType() ValueType {
	offset := binary.LittleEndian.Uint32(i.b) - 4
	tag := binary.LittleEndian.Uint64(i.b[offset:])
	return ValueType(tag & 0XFF)
}

func (i InternalKey) seq() int64 {
	offset := binary.LittleEndian.Uint32(i.b) - 4
	tag := binary.LittleEndian.Uint64(i.b[offset:])
	return int64(tag >> 8)
}
