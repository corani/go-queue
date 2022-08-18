package queue

import "sync"

type Queue[T any] struct {
	mutex sync.RWMutex
	data  []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		mutex: sync.RWMutex{},
		data:  nil,
	}
}

func (q *Queue[T]) Push(val T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.data = append(q.data, val)
}

func (q *Queue[T]) Peek() T {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	return q.peekLocked()
}

//nolint:nonamedreturns
func (q *Queue[T]) peekLocked() (val T) {
	if len(q.data) > 0 {
		return q.data[0]
	}

	// NOTE(daniel): named return is necessary to return the "zero" value for the
	// generic type.
	return
}

func (q *Queue[T]) Pop() T {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return q.popLocked()
}

//nolint:nonamedreturns
func (q *Queue[T]) popLocked() (val T) {
	if len(q.data) > 0 {
		val = q.data[0]

		q.data = q.data[1:]
	}

	// NOTE(daniel): named return is necessary to return the "zero" value for the
	// generic type.
	return val
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) Len() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	return len(q.data)
}
