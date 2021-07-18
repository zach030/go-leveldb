package utils

import (
	"bytes"
)

// a>b 1
// a==b 0
// a<b -1

type Comparator func(a, b interface{}) int

func IntComparator(a, b interface{}) int {
	aInt := a.(int)
	bInt := b.(int)
	switch {
	case aInt > bInt:
		return 1
	case aInt < bInt:
		return -1
	default:
		return 0
	}
}

func UserKeyComparator(a,b interface{})int{
	aK := a.([]byte)
	bK := b.([]byte)
	return bytes.Compare(aK,bK)
}