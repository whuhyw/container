package stack

import (
	"fmt"
	"reflect"
	"container/heap"
)

var _ Stack = (*entry)(nil)

func newEntry() *entry {
	e := &entry{
		s:    []Element{},
		size: 0,
	}
	return e
}

type entry struct {
	s    []Element
	size int
}

func (e *entry) Equal(s Stack) bool {
	reflect.DeepEqual()
	container.
	if e.Size() != s.Size() {
		return false
	}
	return reflect.DeepEqual(e.String(), s.String())
}

func (e *entry) String() string {
	return fmt.Sprint("%#", e.s)
}

func (e *entry) Push(element Element) {
	if e == nil {
		*e = *newEntry()
	}
	e.s = append(e.s, element)
	e.size++
}

func (e *entry) Pop() Element {
	if e.size == 0 {
		return nil
	}
	ret := e.s[e.size-1]
	e.s = e.s[:e.size-1]
	e.size--
	return ret
}

func (e *entry) Top() Element {
	return e.s[e.size-1]
}

func (e *entry) Clear() {
	*e = *newEntry()
}

func (e *entry) Size() int {
	return e.size
}

func (e *entry) Empty() bool {
	return e.size == 0
}
