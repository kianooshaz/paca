package stack

import "errors"

type Stack []string

func New() Stack {
	var stack Stack
	return stack
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) *Stack {
	*s = append(*s, str)
	return s
}

func (s *Stack) Top() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("empty stack")
	}
	index := len(*s) - 1
	str := (*s)[index]
	return str, nil
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("empty stack")
	}
	index := len(*s) - 1
	str := (*s)[index]
	*s = (*s)[:index]
	return str, nil
}
