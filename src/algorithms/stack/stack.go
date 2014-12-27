package stack

import "errors"

type Stack []interface{}

func (stack Stack) Top() (interface{}, error) {
    l := len(stack)
    if l == 0 {
        return nil, errors.New("Stack is empty.")
    }

    return stack[l - 1], nil
}

func (stack *Stack) Push(data interface{}) error {
    if stack == nil {
        return errors.New("Pointer is null.")
    }

    *stack = append(*stack, data)
    return nil
}

func (stack *Stack) Pop() (interface{}, error) {
    l := len(*stack)
    if l == 0 {
        return nil, errors.New("Stack is empty.")
    }
    data := (*stack)[l - 1];
    *stack = (*stack)[: l - 1];
    return data, nil
}

func (stack Stack) Len() int {
    return len(stack)
}
