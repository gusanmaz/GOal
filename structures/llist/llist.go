package llist

import (
	"github.com/gusanmaz/algorithms/helper"
	"github.com/gusanmaz/algorithms/types/printables"
)


type LLisType int
type Direction int

const(
	Single LLisType = iota
	Double
	Circular
)

const(
	Left Direction = iota
	Right
)

type LListProps struct{
	Type LLisType
}

type singleNode struct{
	value printables.Interface
	next  *singleNode
}

type doubleNode struct{
	value printables.Interface
	prev  *doubleNode
	next  *doubleNode
}

type LList struct{
	lListType LLisType
	sHead *singleNode
	len int
}

type LListIterator struct{
	llistType LLisType
	curSNode *singleNode
	curDNode *doubleNode
	direction Direction
}

func New(p LListProps) *LList{
	if p.Type == Single{
		// The compiler performs escape analysis and allocates on the heap when the address escapes the local scope.
		return &LList{Single, nil, 0}
	}else{
		//TODO
		return &LList{}
	}
}


func (l *LList)Insert2Head(val printables.Interface){
	if l.lListType == Single {
		newNode := singleNode{val, l.sHead}
		l.sHead = &newNode
		l.len++
	}else{
		// TODO
	}
}

func (l *LList)Insert2Tail(val printables.Interface){
	if l.lListType == Single {
		tailNode := l.sHead

		newNode := singleNode{val, nil}
		l.len++

		if tailNode == nil {
			l.sHead = &newNode
			return
		}

		for ; tailNode.next != nil; tailNode = tailNode.next {
		}

		tailNode.next = &newNode
	}else{
		// TODO
	}
}

func (l *LList) DeleteHead() printables.Interface{
	if l.lListType == Single {
		head := l.sHead
		if head == nil {
			return nil
		}

		l.sHead = head.next
		l.len--
		return head.value
	}else{
		// TODO
		return nil
	}
}

func (l *LList) DeleteTail() printables.Interface{
	if l.lListType == Single {
		tailNode := l.sHead
		var almostTail *singleNode // nil by default

		if tailNode == nil {
			return nil
		}

		for ; tailNode.next != nil; tailNode = tailNode.next {
			almostTail = tailNode
		}

		if almostTail == nil {
			l.sHead = nil
			l.len--
			return tailNode.value
		}

		almostTail.next = nil
		l.len--
		return tailNode.value
	}else{
		// TODO
		return nil
	}
}

func (l* LList) Iterator() helper.Iterator{
	if l.lListType == Single {
		return &LListIterator{Single,l.sHead, nil, Right}
	}else{
		// TODO
		return  &LListIterator{Double,nil, nil, Right}
	}
}


func (l *LListIterator) Next() interface{}{
	if l.llistType == Single {
		val := l.curSNode.value
		l.curSNode = l.curSNode.next
		return val
	}else{
		//TODO
		return nil
	}
}

func (l *LListIterator) HasNext() bool{
	if l.llistType == Single {
		return l.curSNode != nil
	}else{
		//TODO
		return false
	}
}

