package sudoku

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldGroupString(t *testing.T) {
	size := 2
	s := New(size)
	f := NewField(s, 0, 4)
	fg := NewFieldGroup(s, size, "col 1")
	assert.Equal(t, "", fg.String())
	fg.Fields[0] = &f
	assert.Equal(t, "4", fg.String())
}

func TestFieldGroupIsSolved(t *testing.T) {
	size := 2
	s := New(size)
	f := NewField(s, 0, 2)
	f2 := NewField(s, 0, 0)
	fg := NewFieldGroup(s, size, "col 1")
	fg.Fields[0] = &f
	fg.Fields[1] = &f2
	assert.False(t, fg.IsSolved())
	f2.Value = 1
	assert.True(t, fg.IsSolved())
}

func TestFieldGroupSolve(t *testing.T) {
	s, _ := FromReader(strings.NewReader(`-- 34 ---- ---- ----`))
	f1 := s.Fields[0]
	f2 := s.Fields[1]
	fg := s.GetRow(f1)
	res := fg.Solve()
	assert.False(t, res.FoundNew)
	f2.Value = 2
	res = fg.Solve()
	assert.True(t, res.FoundNew)
	assert.Equal(t, 1, f1.Value)
	assert.True(t, fg.IsSolved())
	res = fg.Solve()
	assert.False(t, res.FoundNew)
}
