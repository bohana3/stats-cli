package stats

import (
	"errors"
	"sync"
)

type Queue struct {
	mu sync.Mutex
	q []string
	capacity int
}

// FixSizeQueue (FIFO)
type FixSizeQueue interface {
	Insert()
	Remove()
}

// Insert inserts the item into the queue
func (q *Queue) Insert(item string) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.q = append(q.q, item)

	for len(q.q) > int(q.capacity) {
		q.q = q.q[1:]
	}

	return nil
}

// Remove removes the oldest element from the queue
func (q *Queue) Remove() (string, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) > 0 {
		item := q.q[0]
		q.q = q.q[1:]
		return item, nil
	}
	return "", errors.New("Queue is empty")
}

// CreateQueue creates an empty queue with desired capacity
func CreateQueue(capacity int) *Queue {
	return &Queue{
		capacity: capacity,
		q:        make([]string, 0, capacity),
	}
}