package stack

import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack *Stack) Push(value interface{}) *Stack {
	*stack = append(*stack, value)
	return stack
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Can't Top() an empty stack")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	self := *stack
	if len(self) == 0 {
		return nil, errors.New("Can't Pop() an empty stack")
	}
	var x = self[len(self)-1]
	*stack = self[:len(self)-1]
	return x, nil
}
