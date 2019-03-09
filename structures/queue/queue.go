package queue

import (
	"github.com/gusanmaz/algorithms/types/printables"
)

type QueueType int

const(
	Array = iota
	List
	Circular
)

type QueueProps struct{
	QType QueueType
	Size int
}

type ArrayQueue struct{
	elems []printables.Interface
	front int
	rear  int
	cnt   int
}

type Queue struct{
	qType QueueType
	arrQueue *ArrayQueue
}

func New(props QueueProps) *Queue{
	if props.QType == List{
		//TODO
	}else if props.QType == Circular{
		//TODO
	}else{
		elems := make([]printables.Interface, 10)
		arrQueue := ArrayQueue{elems, 0, 0, 0}
		return &Queue{Array, &arrQueue}
	}
	return nil
}

func arrayEnqueue(queue *ArrayQueue, elem printables.Interface){
	if elem == nil{
		return
	}

	capacity := cap(queue.elems)

	queue.elems[queue.rear] = elem
	queue.cnt++
	queue.rear = (queue.rear + 1) % capacity
}

func arrayDequeue(queue *ArrayQueue) (printables.Interface) {
	capacity := cap(queue.elems)
	if queue.cnt == 0{
		return nil
	}

	elem := queue.elems[queue.front]
	queue.elems[queue.front] = nil
	queue.cnt--
	queue.front =  (queue.front + 1) % capacity

	return elem
}

func (queue *Queue) Enqueue(elem printables.Interface){
	if queue.qType == Array{
		arr := queue.arrQueue.elems
		capacity := cap(arr)
		if capacity == queue.arrQueue.cnt{
			elems := make([]printables.Interface, 2 * capacity)
			biggerQueue := ArrayQueue{elems, 0, 0, 0}
			for elem := arrayDequeue(queue.arrQueue); elem != nil; elem = arrayDequeue(queue.arrQueue) {
				arrayEnqueue(&biggerQueue, elem)
			}
			queue.arrQueue = &biggerQueue
		}
		arrayEnqueue(queue.arrQueue, elem)
	}
	//TODO
}

func (queue *Queue) Dequeue() printables.Interface{
	if queue.qType == Array{
		arr := queue.arrQueue.elems
		capacity := cap(arr)
		if capacity > 4 * queue.arrQueue.cnt{
			elems := make([]printables.Interface, capacity / 2)
			smallerQueue := ArrayQueue{elems, 0, 0, 0}
			for elem := arrayDequeue(queue.arrQueue); elem != nil; elem = arrayDequeue(queue.arrQueue) {
				arrayEnqueue(&smallerQueue, elem)
			}
			queue.arrQueue = &smallerQueue
		}
		elem := arrayDequeue(queue.arrQueue)
		return elem
	}
	//TODO
	return nil
}


