package datastructure

//StackInt data structure
type StackRune interface {
	 Pop() rune
	 Top() rune
	 Push(value rune)
	 Empty() bool
	 Size() int
}

type StackRuneImp struct {
	list []rune
}

func NewStackRune() StackRune {
	return &StackRuneImp{ list: []rune{}}
}

func (stack *StackRuneImp) Size() int {
	return len(stack.list)
}

func (stack *StackRuneImp) Empty() bool {

	return stack.Size() == 0
}

func (stack *StackRuneImp) Push(value rune) {
	stack.list = append(stack.list, value)
}

func (stack *StackRuneImp) Top() rune {
	return stack.list[stack.Size() - 1]
}

func (stack *StackRuneImp) Pop() rune {
	result := stack.Top()
	stack.list = stack.list[: stack.Size() -1]
	return result
}