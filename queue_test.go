package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(3, q.Len())

	assert.Equal(1, q.Dequeue())
	assert.Equal(2, q.Dequeue())
	assert.Equal(3, q.Dequeue())
}
