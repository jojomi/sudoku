package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldPossibleValues(t *testing.T) {
	size := 2
	s := New(size)
	f := NewField(s, 0, 0)
	assert.Equal(t, size*size, len(f.PossibleValues()))

	f.DenyValue(1)
	assert.Equal(t, size*size-1, len(f.PossibleValues()))
}

func TestFieldSolvable(t *testing.T) {
	size := 2
	s := New(size)
	f := NewField(s, 0, 0)

	assert.False(t, f.Solvable())
	f.DenyValue(1)
	assert.False(t, f.Solvable())
	f.DenyValue(2)
	assert.False(t, f.Solvable())
	f.DenyValue(4)
	assert.True(t, f.Solvable())
}

func TestFieldIsSolved(t *testing.T) {
	size := 2
	s := New(size)

	f := NewField(s, 0, 1)
	assert.True(t, f.IsSolved())

	f2 := NewField(s, 0, 0)
	assert.False(t, f2.IsSolved())
	f2.Value = 2
	assert.True(t, f2.IsSolved())
}

func TestFieldSolve(t *testing.T) {
	size := 2
	s := New(size)

	f := NewField(s, 0, 0)
	assert.False(t, f.Solve())
	f.DenyValue(2)
	assert.False(t, f.Solve())
	f.DenyValue(3)
	assert.False(t, f.Solve())
	assert.Equal(t, 0, f.Value)
	f.DenyValue(4)
	assert.True(t, f.Solve())
	assert.Equal(t, 1, f.Value)
}

func TestFieldString(t *testing.T) {
	size := 2
	s := New(size)
	f := NewField(s, 0, 0)
	assert.Equal(t, ".", f.String())

	f.Value = 2
	assert.Equal(t, "2", f.String())
}
