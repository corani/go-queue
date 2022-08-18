package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		q := New[int]()

		assert.NotNil(t, q)
		assert.Equal(t, 0, q.Len())
		assert.True(t, q.Empty())

		assert.Equal(t, 0, q.Peek())
		assert.Equal(t, 0, q.Pop())
	})

	t.Run("push", func(t *testing.T) {
		q := New[int]()

		assert.NotNil(t, q)
		q.Push(33)

		assert.Equal(t, 1, q.Len())
		assert.False(t, q.Empty())

		assert.Equal(t, 33, q.Peek())
		assert.Equal(t, 33, q.Pop())
		assert.True(t, q.Empty())
	})
}
