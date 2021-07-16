package skiplist

import (
	"fmt"
	"go-leveldb/utils"
	"math/rand"
	"testing"
)

func Test_Insert(t *testing.T) {
	skipList := New(utils.IntComparator)
	for i := 0; i < 10; i++ {
		skipList.Insert(rand.Int() % 10)
	}
	it := skipList.NewIterator()
	for it.SeekToFirst(); it.Valid(); it.Next() {
		fmt.Println(it.Key())
	}
	fmt.Println()
	for it.SeekToLast(); it.Valid(); it.Prev() {
		fmt.Println(it.Key())
	}

}
