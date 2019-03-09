package structures

import (
	"github.com/gusanmaz/algorithms/helper"
	"github.com/gusanmaz/algorithms/types/printables"
)

// One direction linked list implementation

type uniNode struct{
	value printables.Interface
	next  *uniNode
}

type UniList struct{
	head *uniNode
	len  int
}

type UniListIterator struct{
	curNode *uniNode
}

func NewUniList() *UniList{
	// The compiler performs escape analysis and allocates on the heap when the address escapes the local scope.
	return &UniList{nil, 0}
}

func (u *UniList)Insert2Head(val printables.Interface){
	newNode := uniNode{val, u.head}
	u.head   =  &newNode
	u.len++
}

func (u *UniList)Insert2Tail(val printables.Interface){
	tailNode := u.head

	newNode := uniNode{val, nil}
	u.len++

	if tailNode == nil{
		u.head = &newNode
		return
	}

	//var almostTail *uniNode // nil by default
	for ;tailNode.next != nil; tailNode = tailNode.next{
		//almostTail = tailNode
	}

	//if almostTail == nil{
	//	u.head.next = &newNode
	//	return
	//}

	tailNode.next = &newNode
}

func (u *UniList) DeleteHead() printables.Interface{
	head := u.head
	if head == nil{
		return nil
	}

	u.head = head.next
	u.len--
	return head.value
}

func (u *UniList) DeleteTail() printables.Interface{
	tailNode := u.head
	var almostTail *uniNode // nil by default

	if tailNode == nil{
		return nil
	}

	for ;tailNode.next != nil; tailNode = tailNode.next{
		almostTail = tailNode
	}

	if almostTail == nil{
		u.head = nil
		u.len--
		return tailNode.value
	}

	almostTail.next = nil
	u.len--
	return tailNode.value
}

func (u* UniList) Iterator() helper.Iterator{
	return &UniListIterator{u.head}
}


func (u *UniListIterator) Next() interface{}{
	val := u.curNode.value
	u.curNode = u.curNode.next
	return val
}

func (u *UniListIterator) HasNext() bool{
	return u.curNode != nil
}







