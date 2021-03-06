package stack

type Element interface{}

type Stack interface {
	Push(e Element)
	Pop() Element
	Top() Element
	Clear()
	Size() int
	Empty() bool
	String() string
	Equal(s Stack) bool
}

func NewStack() *entry {
	return newEntry()
}

func NewThreadSafeStack() *threadSafeStack {
	return newThreadSafeStack()
}
