package goQueue

import "container/list"

type secureQueue struct {
	data *list.List
	size int
	lock chan int8
}

func NewSecureQueue() *secureQueue {
	q := new(secureQueue)
	q.init()
	return q
}

func (q *secureQueue) init() {
	q.data = list.New()
	q.lock = make(chan int8, 1)
}

func (q *secureQueue) Size() int {
	q.lock <- 1
	defer func(lock chan int8) {<- lock}(q.lock)
	return q.size
}

func (q *secureQueue) Empty() bool {
	q.lock <- 1
	defer func(lock chan int8) {<- lock}(q.lock)
	return q.size == 0
}

func (q *secureQueue) Front() interface{} {
	q.lock <- 1
	defer func(lock chan int8) {<- lock}(q.lock)
	return q.data.Front().Value
}

func (q *secureQueue) Back() interface{} {
	q.lock <- 1
	defer func(lock chan int8) {<- lock}(q.lock)
	return q.data.Back().Value
}

func (q *secureQueue) Push(value interface{}) {
	q.lock <- 1
	defer func(lock chan int8) {<- lock}(q.lock)
	q.data.PushBack(value)
	q.size++
}

func (q *secureQueue) Pop() {
	q.lock <- 1
	defer func(lock chan int8) {<- lock}(q.lock)
	if q.size > 0 {
		tmp := q.data.Back()
		q.data.Remove(tmp)
		q.size--
	}
}