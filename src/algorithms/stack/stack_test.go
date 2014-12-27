package stack

import "testing"

var data = []int { 1, 2, 3, 4, 5 }
var dataPopped = []int { 5, 4, 3, 2, 1 }

func TestPush(t *testing.T) {
    stack := Stack {}
    for i, l := 0, len(data); i < l; i++ {
        stack.Push(data[i])
        if stack[i] != data[i] {
            t.Error("For", i, "Actual", stack[i], "Expected", data[i])
        }
    }
}

func TestLen(t *testing.T) {
    stack := Stack {}
    length := stack.Len()
    if length != 0 {
        t.Error("For", "Len()", "Actual", length, "Expected", 0)
        return
    }

    for i, l := 0, len(data); i < l; i++ {
        stack.Push(data[i])
        length = stack.Len()
        expected := i + 1
        if length != expected {
            t.Error("For", "Len()", "Actual", length, "Expected", expected)
            return
        }
    }
}

func TestPop(t *testing.T) {
    stack := Stack {}
    for i, l := 0, len(data); i < l; i++ {
        stack.Push(data[i])
    }

    for i, l := 0, len(dataPopped); i < l; i++ {
        d, _ := stack.Pop()
        if d != dataPopped[i] {
            t.Error("For", "Pop()", "Actual", d, "Expected", dataPopped[i])
            return
        }
    }
}

func TestTop(t *testing.T) {
    stack := Stack {}
    for i, l := 0, len(data); i < l; i++ {
        stack.Push(data[i])
        d, _ := stack.Top()
        if d != data[i] {
            t.Error("For", "Top()", "Actual", d, "Expected", data[i])
            return
        }
    }
}
