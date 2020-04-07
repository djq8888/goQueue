package goQueue

import "container/list"

type Queue struct {
	data *list.List
	size int
}

func NewQueue() *Queue {
	q := new(Queue)
	q.init()
	return q
}

func (q *Queue) init() {
	q.data = list.New()
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Empty() bool {
	return q.size == 0
}

func (q *Queue) Front() interface{} {
	return q.data.Front().Value
}

func (q *Queue) Back() interface{} {
	return q.data.Back().Value
}

func (q *Queue) Push(value interface{}) {
	q.data.PushBack(value)
	q.size++
}

func (q *Queue) Pop() {
	if q.size > 0 {
		tmp := q.data.Back()
		q.data.Remove(tmp)
		q.size--
	}
}