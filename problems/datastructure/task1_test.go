package problems

import (
	"github.com/alexrv11/algorithms/datastructure"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidateCase1(t *testing.T) {
	text := "{[()]}"
	task := &Problem1{stack: datastructure.NewStackRune(), text: text}

	result := task.Validate()

	assert.True(t, result)
}

func TestValidateCase2(t *testing.T) {
	text := "{[]}{()}"
	task := &Problem1{stack: datastructure.NewStackRune(), text: text}

	result := task.Validate()

	assert.True(t, result)
}

func TestValidateCase3(t *testing.T) {
	text := "{[]}]()"
	task := &Problem1{stack: datastructure.NewStackRune(), text: text}

	result := task.Validate()

	assert.False(t, result)
}

func TestValidateCase4(t *testing.T) {
	text := "{{{[["
	task := &Problem1{stack: datastructure.NewStackRune(), text: text}

	result := task.Validate()

	assert.False(t, result)
}

func TestValidateCase5(t *testing.T) {
	text := "]]]}}"
	task := &Problem1{stack: datastructure.NewStackRune(), text: text}

	result := task.Validate()

	assert.False(t, result)
}
