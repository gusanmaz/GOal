package main

import (
	"fmt"
	"github.com/gusanmaz/algorithms/algorithms/search"
	"github.com/gusanmaz/algorithms/algorithms/sort"
	"github.com/gusanmaz/algorithms/generators"
	"github.com/gusanmaz/algorithms/structures/bag"
	"github.com/gusanmaz/algorithms/structures/disjointSet"
	"github.com/gusanmaz/algorithms/structures/llist"
	"github.com/gusanmaz/algorithms/structures/queue"
	"github.com/gusanmaz/algorithms/structures/stack"
	"github.com/gusanmaz/algorithms/types/comparables"
	"github.com/gusanmaz/algorithms/types/printables"
)

func main() {
	algIntArr := comparables.IntArr([]int{5, 6, 7, 8, 9})
	algIntArr3 := comparables.IntArr([]int{})
	algIntArr2 := comparables.IntArr([]int{50, 6, 72, 81, 9})
	Float64Arr := comparables.Float64Arr([]float64{5.1, 6.0, 7.0, 8.0, 9.0})

	fmt.Println(search.Binary(algIntArr, comparables.Int(6)))
	fmt.Println(search.Binary(Float64Arr, comparables.Float64(7.0)))
	fmt.Println(search.Linear(algIntArr, comparables.Int(6)))

	sort.Bubble(algIntArr2)
	fmt.Println(algIntArr2)

	sort.Selection(algIntArr2)
	fmt.Println(algIntArr2)
	sort.Selection(algIntArr3)
	fmt.Println(algIntArr3)

	bag := bag.New()
	bag.Add(printables.Int(122))
	bag.Add(printables.Int(142))
	bag.Add(printables.Int(152))
	bag.Add(printables.Int(127))
	fmt.Println(bag)
	bit := bag.Iterator()
	for bit.HasNext() {
		fmt.Println(bit.Next())
	}

	kk := generators.IntArr(&generators.IntArrProps{-10, 11, 40})
	fmt.Println(kk)

	ll := generators.IntArr(&generators.IntArrProps{Size: 20000})
	fmt.Println(ll)

	mm := generators.SFloat64Arr(&generators.SFloat64ArrProps{})
	fmt.Println(mm)

	stack := stack.New()
	stack.Push(printables.Int(6))
	stack.Push(printables.Int(7))
	stack.Push(printables.Int(8))

	sIt := stack.Iterator()

	fmt.Println(stack)

	for sIt.HasNext() {
		fmt.Println("It", sIt.Next())
	}

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println("Deneme")

	llist := llist.New(llist.LListProps{llist.Single})
	llist.Insert2Tail(printables.Int(6))
	llist.DeleteHead()
	llist.Insert2Head(printables.Int(67))
	llist.Insert2Head(printables.Int(64))
	llist.Insert2Head(printables.Int(63))
	llist.Insert2Tail(printables.Int(7))
	llist.DeleteTail()



	/*llist2 := llist.New(llist.LListProps{llist.Single})
	llist2.Insert2Tail(printables.Int(10))
	llist2.Insert2Tail(printables.Int(15))
	llist2.Insert2Tail(printables.Int(20))
	llist2.Insert2Tail(printables.Int(30))

	unit := llist.Iterator()

	for unit.HasNext() {
		fmt.Println(unit.Next())
	}

	unit2 := llist2.Iterator()
	fmt.Println("hello")
	for unit2.HasNext() {
		fmt.Println(unit2.Next())
	}*/

	q := queue.New(queue.QueueProps{})
	q.Enqueue(printables.Int(1))
	q.Enqueue(printables.Int(2))
	q.Enqueue(printables.Int(3))
	q.Dequeue()
	q.Enqueue(printables.Int(4))
	q.Enqueue(printables.Int(5))
	q.Enqueue(printables.Int(6))
	//q.Dequeue()
	q.Enqueue(printables.Int(7))
	q.Enqueue(printables.Int(8))
	q.Enqueue(printables.Int(9))
	q.Enqueue(printables.Int(10))
	q.Enqueue(printables.Int(11))
	q.Enqueue(printables.Int(12))
	q.Enqueue(printables.Int(13))

	dset := disjointSet.New(disjointSet.Props{disjointSet.FastFind, 20})
	dset.Union(5,6)
	dset.Union(7,8)
	dset.Union(1,6)
	fmt.Println(dset.Find(1), dset.Find(5))





}
