package data_structures

type Queue[TValue any] []TValue

func (queue *Queue[TValue]) IsEmpty() bool {
	return len(*queue) == 0
}

func (queue *Queue[TValue]) Enqueue(value TValue) {
	*queue = append(*queue, value)
}

func (queue *Queue[TValue]) Dequeue() TValue {
	if queue.IsEmpty() {
		panic("queue is empty")
	}

	value := (*queue)[0]
	*queue = (*queue)[1:]

	return value
}

func (queue *Queue[TValue]) Peek() TValue {
	if queue.IsEmpty() {
		panic("queue is empty")
	}

	return (*queue)[0]
}

func (queue *Queue[TValue]) Size() int {
	return len(*queue)
}

func (queue *Queue[TValue]) Clear() {
	*queue = (*queue)[:0]
}
