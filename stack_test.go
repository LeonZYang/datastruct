package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	assert := assert.New(t)
	q := NewStack()

	q.Push(1)
	q.Push(2)
	q.Push(3)
	assert.Equal(3, q.Len())

	assert.Equal(3, q.Pop())
	assert.Equal(2, q.Pop())
	assert.Equal(1, q.Pop())
}
