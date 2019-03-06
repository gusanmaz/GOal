package structures

import (
	"github.com/gusanmaz/algorithms/types/printables"
	"github.com/gusanmaz/algorithms/helper"

)
import _ "fmt"

type Bag struct{
	items []printables.Interface
}

type BagIterator struct{
	items []printables.Interface
	len   int
	curInd int
}

func (b *BagIterator) Next() interface{}{
	b.curInd++
	return b.items[b.curInd - 1]
}

func (b *BagIterator) HasNext() bool{
	return b.len > b.curInd
}

func NewBag() *Bag{
	var b Bag = Bag{}
	b.items = make([]printables.Interface, 10)
	b.items = b.items[:0]
	return &b
}


func (b *Bag) Add(elem printables.Interface){
	bagLen   := len(b.items)
	bagCap   := cap(b.items)
	if bagLen == bagCap {
		doubleArr := make([]printables.Interface, bagCap * 2)
		copy(doubleArr, b.items)
		b.items = doubleArr
	}
	b.items = b.items[:(bagLen + 1)]
	b.items[bagLen] = elem
}

func (b* Bag) String() string{
	bagLen   := len(b.items)
	elems := ""
	for i := 0; i < bagLen; i++{
		elems = elems + b.items[i].String() + ","
	}
	if bagLen == 0{
		return elems
	}
	return elems[:len(elems) - 1]
}

func (b* Bag) Iterator() helper.Iterator{
	bagLen   := len(b.items)
	it := BagIterator{}
	it.items = b.items[:bagLen]
	it.curInd = 0
	return &it
}


