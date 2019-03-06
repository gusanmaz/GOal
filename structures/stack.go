package structures

import (
	"github.com/gusanmaz/algorithms/helper"
	"github.com/gusanmaz/algorithms/types/printables"
)
import _ "fmt"

type Stack struct{
	items []printables.Interface
}

type StackIterator struct{
	items []printables.Interface
	curInd int
}

func (b *StackIterator) Next() interface{}{
	b.curInd--
	return b.items[b.curInd]
}

func (b *StackIterator) HasNext() bool{
	return 0 < b.curInd
}

func NewStack() *Stack{
	var b Stack = Stack{}
	b.items = make([]printables.Interface, 10)
	b.items = b.items[:0]
	return &b
}

func (s* Stack) Iterator() helper.Iterator{
	return &StackIterator{s.items,len(s.items)}
}


func (b *Stack) Push(elem printables.Interface){
	stackLen := len(b.items)
	stackCap := cap(b.items)
	if stackLen == stackCap{
		doubleArr := make([]printables.Interface, stackCap * 2)
		copy(doubleArr, b.items)
		b.items = doubleArr
	}
	b.items = b.items[:(stackLen + 1)]
	b.items[stackLen] = elem
}

func (b *Stack) Pop() printables.Interface{
	stackLen := len(b.items)
	stackCap := cap(b.items)
	if stackLen == 0 {
		panic("Attempt to pop from a empty stack")
	}
	stackLen--
	popElem := b.items[stackLen]
	b.items[stackLen] = nil // Avoid loitering

	if stackCap >= 4 * stackLen{
		smallerArr := make([]printables.Interface, stackCap / 2)
		copy(smallerArr, b.items)
		b.items = smallerArr[:stackLen]
	}
	return popElem
}

func (b* Stack) String() string{
	elems := ""
	stackLen := len(b.items)
	for i := 0; i < stackLen; i++{
		elems = elems + b.items[i].String() + ","
	}
	if stackLen == 0{
		return elems
	}
	elems = elems[:len(elems) - 1]
	return elems
}


