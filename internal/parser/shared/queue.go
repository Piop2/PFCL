package shared

// Queue is a generic FIFO queue.
type Queue[T any] struct {
	data []T
}

// Enqueue adds an item to the end of the queue.
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

// Dequeue removes and returns the first item from the queue.
// If the queue is empty, the second return value is false.
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.data) == 0 {
		return zero, false
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v, true
}

// Peek returns the first item without removing it.
// If the queue is empty, the second return value is false.
func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if len(q.data) == 0 {
		return zero, false
	}
	return q.data[0], true
}

// Len returns the number of items in the queue.
func (q *Queue[T]) Len() int {
	return len(q.data)
}

// IsEmpty returns true if the queue has no items.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}
