package stack

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

func (s *StackIterator) Next() interface{}{
	s.curInd--
	return s.items[s.curInd]
}

func (s *StackIterator) HasNext() bool{
	return 0 < s.curInd
}

func New() *Stack{
	var b Stack = Stack{}
	b.items = make([]printables.Interface, 10)
	b.items = b.items[:0]
	return &b
}

func (s * Stack) Iterator() helper.Iterator{
	return &StackIterator{s.items,len(s.items)}
}


func (s *Stack) Push(elem printables.Interface){
	stackLen := len(s.items)
	stackCap := cap(s.items)
	if stackLen == stackCap{
		doubleArr := make([]printables.Interface, stackCap * 2)
		copy(doubleArr, s.items)
		s.items = doubleArr
	}
	s.items = s.items[:(stackLen + 1)]
	s.items[stackLen] = elem
}

func (s *Stack) Pop() printables.Interface{
	stackLen := len(s.items)
	stackCap := cap(s.items)
	if stackLen == 0 {
		panic("Attempt to pop from an empty stack")
	}
	stackLen--
	popElem := s.items[stackLen]
	s.items[stackLen] = nil // Avoid loitering

	if stackCap >= 4 * stackLen{
		smallerArr := make([]printables.Interface, stackCap / 2)
		copy(smallerArr, s.items)
		s.items = smallerArr[:stackLen]
	}
	return popElem
}

func (s * Stack) String() string{
	elems := ""
	stackLen := len(s.items)
	for i := 0; i < stackLen; i++{
		elems = elems + s.items[i].String() + ","
	}
	if stackLen == 0{
		return elems
	}
	elems = elems[:len(elems) - 1]
	return elems
}


