package datastructure

import(
	"testing"
  "github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	stack := NewStackRune()

	assert.NotNil(t, stack)
}

func TestSize(t *testing.T) {
	stack := NewStackRune()

	size := stack.Size()

	assert.Equal(t, 0, size)
}

func TestPush(t *testing.T) {
	stack := NewStackRune()

	stack.Push('a')

	assert.Equal(t, 1, stack.Size())
}

func TestTop(t *testing.T) {
	stack := NewStackRune()
	stack.Push('a')
	stack.Push('q')
	value := stack.Top()

	assert.Equal(t, 'q', value)
}

func TestPop(t *testing.T) {
	stack := NewStackRune()
	stack.Push('a')
	stack.Push('q')

	deletedValue := stack.Pop()
	value := stack.Top()

	assert.Equal(t, 'q', deletedValue)
	assert.Equal(t, 'a', value)
}