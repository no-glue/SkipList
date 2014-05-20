// SkipList
package main

import (
	"fmt"
	"math/rand"
)

const MAX_LEVEL = 32

type Ordered interface {
	Less(other Ordered) bool
	LessEquel(other Ordered) bool
	Equel(other Ordered) bool
}

type SkipListNode struct {
	Key     Ordered
	Value   interface{}
	Forward []*SkipListNode
}

func NewNode(level int, key Ordered, value interface{}) *SkipListNode {
	node := &SkipListNode{
		Key:     key,
		Value:   value,
		Forward: make([]*SkipListNode, level),
	}
	for i := 0; i < level; i++ {
		node.Forward[i] = nil
	}
	return node
}

type SkipList struct {
	Level  int
	Head   *SkipListNode
	length int
}

func NewSkipList() *SkipList {
	sl := &SkipList{
		Level:  0,
		Head:   NewNode(MAX_LEVEL, nil, "HEAD"),
		length: 0,
	}
	return sl
}

func RandomLevel() int {
	level := 1
	for rand.Intn(2) == 0 && level < MAX_LEVEL {
		level++
	}
	return level
}

func (sl *SkipList) Insert(key Ordered, value interface{}) bool {
	update := make([]*SkipListNode, MAX_LEVEL)
	p, q := sl.Head, sl.Head
	level := sl.Level
	for {
		q = p.Forward[level]
		for q != nil && q.Key.Less(key) {
			p = q
			q = p.Forward[level]
		}
		update[level] = p
		level--
		if level < 0 {
			break
		}
	}
	if q != nil && q.Key.Equel(key) {
		q.Value = value
		return false
	}
	k := RandomLevel()
	if k > sl.Level {
		sl.Level++
		k = sl.Level
		update[k-1] = sl.Head
	}
	q = NewNode(k, key, value)
	for i := k - 1; i >= 0; i-- {
		p = update[i]
		q.Forward[i] = p.Forward[i]
		p.Forward[i] = q
	}
	sl.length++
	return true
}

func (sl *SkipList) Delete(key Ordered) bool {
	update := make([]*SkipListNode, MAX_LEVEL)
	p, q := sl.Head, sl.Head
	level := sl.Level
	for {
		q = p.Forward[level]
		for q != nil && q.Key.Less(key) {
			p = q
			q = p.Forward[level]
		}
		update[level] = p
		level--
		if level < 0 {
			break
		}
	}
	if q != nil && q.Key.Equel(key) {
		for i := 0; i < sl.Level; i++ {
			if update[i].Forward[i] == q {
				update[i].Forward[i] = q.Forward[i]
			}
		}
		q = nil
		for i := sl.Level - 1; i >= 0; i-- {
			if sl.Head.Forward[i] == nil {
				sl.Level--
			}
		}
		sl.length--
		return true
	}
	return false
}

func (sl *SkipList) Find(key Ordered) interface{} {
	p, q := sl.Head, sl.Head
	for i := sl.Level - 1; i >= 0; i-- {
		q = p.Forward[i]
		for q != nil && q.Key.LessEquel(key) {
			if q.Key.Equel(key) {
				return q.Value
			}
			p = q
			q = p.Forward[i]
		}
	}
	return -1
}

func (sl *SkipList) Len() int {
	return sl.length
}

func (sl *SkipList) Keys() []Ordered {
	p, q := sl.Head, sl.Head
	keys := []Ordered{}
	p = sl.Head
	q = p.Forward[0]
	for q != nil {
		keys = append(keys, q.Key)
		p = q
		q = p.Forward[0]
	}
	return keys
}

func (sl *SkipList) Print() {
	p, q := sl.Head, sl.Head
	for i := sl.Level - 1; i >= 0; i-- {
		p = sl.Head
		q = p.Forward[i]
		for q != nil {
			fmt.Printf("<%#v, %#v>", q.Key, q.Value)
			p = q
			q = p.Forward[i]
		}
		fmt.Println("\n")
	}
	fmt.Println("\n")
}
