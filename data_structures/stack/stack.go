package data_structures

type Stack[TValue any] []TValue

// Check if the stack is empty
func (stack *Stack[TValue]) IsEmpty() bool {
	return len(*stack) == 0
}

// Push an element onto the top of the stack
func (stack *Stack[TValue]) Push(value TValue) {
	*stack = append(*stack, value)
}

// Pop an element off the top of the stack
func (stack *Stack[TValue]) Pop() TValue {
	if stack.IsEmpty() {
		panic("stack is empty")
	}

	value := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	return value
}

// Peek the top element of the stack
func (stack *Stack[TValue]) Peek() TValue {
	if stack.IsEmpty() {
		panic("stack is empty")
	}

	return (*stack)[len(*stack)-1]
}

// Get the size of the stack
func (stack *Stack[TValue]) Size() int {
	return len(*stack)
}

// Clear the stack
func (stack *Stack[TValue]) Clear() {
	*stack = (*stack)[:0]
}
