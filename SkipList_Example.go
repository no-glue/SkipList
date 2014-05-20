// main
package main

import (
	"fmt"
)

type KeyType int32

func (kt KeyType) Less(other Ordered) bool {
	return kt < other.(KeyType)
}
func (kt KeyType) LessEquel(other Ordered) bool {
	return kt <= other.(KeyType)
}
func (kt KeyType) Equel(other Ordered) bool {
	return kt == other.(KeyType)
}

func main() {
	sl := NewSkipList()
	for i := 20; i >= 0; i-- {
		sl.Insert(KeyType(i), i)
	}
	sl.Insert(KeyType(16), "heool")
	sl.Print()

	sl.Delete(KeyType(4))
	sl.Delete(KeyType(9))
	sl.Delete(KeyType(14))

	sl.Print()

	for i := 20; i >= 0; i-- {
		if sl.Find(KeyType(i)) != i {
			fmt.Println(i)
		}
	}

	fmt.Println("len:", sl.Len())
	fmt.Println("Keys:", sl.Keys())
}
