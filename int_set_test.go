package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	s := NewIntSet()

	assert.Equal(t, "[]", s.String())
	s.Add(3)
	assert.Equal(t, "[3]", s.String())
}

func TestAdd(t *testing.T) {
	s := NewIntSet()
	assert.Len(t, s.Values(), 0)
	s.Add(3)
	assert.Len(t, s.Values(), 1)
	s.Add(3)
	assert.Len(t, s.Values(), 1)
}

func TestContains(t *testing.T) {
	s := NewIntSet()
	s.Add(3)
	assert.False(t, s.Contains(2))
	assert.True(t, s.Contains(3))
}

func TestValues(t *testing.T) {
	s := NewIntSet()
	assert.Len(t, s.Values(), 0)
	s.Add(3)
	assert.Len(t, s.Values(), 1)
}
