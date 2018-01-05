package sudoku

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromFile(t *testing.T) {
	s, err := FromFile("testfiles/small.sudoku")

	assert.Nil(t, err)
	assert.Equal(t, 2, s.Size)
	assert.Equal(t, 16, s.SolvedFieldCount())
}

func TestFromReader(t *testing.T) {
	s, err := FromReader(strings.NewReader(`
		12 34
		34 21

		21 -3
		43 12
	`))

	assert.Nil(t, err)
	assert.Equal(t, 2, s.Size)
	assert.Equal(t, 15, s.SolvedFieldCount())
}

func TestSudokuString(t *testing.T) {
	s, _ := FromFile("testfiles/small.sudoku")

	assert.Equal(t, "+--+--+\n|12|34|\n|34|21|\n+--+--+\n|21|43|\n|43|12|\n+--+--+\n", s.String())
}

func TestIsSolved(t *testing.T) {
	s, _ := FromReader(strings.NewReader(`
		12 34
		34 21

		21 -3
		43 12
	`))

	assert.False(t, s.IsSolved())

	s2, _ := FromReader(strings.NewReader(`
		12 34
		34 21

		21 43
		43 12
	`))

	assert.True(t, s2.IsSolved())
}

func TestSolveStepRow(t *testing.T) {
	s, _ := FromReader(strings.NewReader(`12 -- ---- ---- ----`))
	assert.Equal(t, 2, s.SolvedFieldCount())
	s.SolveStep(SolveOptions{})
	assert.Equal(t, 2, s.SolvedFieldCount())

	s2, _ := FromReader(strings.NewReader(`12 4- ---- ---- ----`))
	assert.Equal(t, 3, s2.SolvedFieldCount())
	s2.SolveStep(SolveOptions{})
	assert.Equal(t, 4, s2.SolvedFieldCount())
}

func TestSolveStepCol(t *testing.T) {
	s, _ := FromReader(strings.NewReader(`2- 1- ---- ---- ----`))
	assert.Equal(t, 2, s.SolvedFieldCount())
	s.SolveStep(SolveOptions{})
	assert.Equal(t, 2, s.SolvedFieldCount())

	s2, _ := FromReader(strings.NewReader(`4- 3- -- 1- ---- ----`))
	assert.Equal(t, 3, s2.SolvedFieldCount())
	s2.SolveStep(SolveOptions{})
	assert.Equal(t, 4, s2.SolvedFieldCount())
}

func TestSolveStepBlock(t *testing.T) {
	s, _ := FromReader(strings.NewReader(`2- 1- ---- ---- ----`))
	assert.Equal(t, 2, s.SolvedFieldCount())
	s.SolveStep(SolveOptions{})
	assert.Equal(t, 2, s.SolvedFieldCount())

	s2, _ := FromReader(strings.NewReader(`42 3- -- -- ---- ----`))
	assert.Equal(t, 3, s2.SolvedFieldCount())
	s2.SolveStep(SolveOptions{})
	assert.Equal(t, 4, s2.SolvedFieldCount())
}

func TestSolveStepNonValues(t *testing.T) {
	// it's about the top-left corner that can be deduced
	s, _ := FromReader(strings.NewReader(`
		-- 1-
		-2 --

		41 --
		-- --
	`))
	s.SolveStep(SolveOptions{})
	assert.True(t, s.Fields[0].IsSolved())
	assert.Equal(t, 3, s.Fields[0].Value)
}

func TestSolveBrute(t *testing.T) {
	s, _ := FromReader(strings.NewReader(`
		12 34
		34 ..

		21 ..
		43 ..
	`))
	s.Solve(SolveOptions{DontDeduce: true})
}

func TestSolveIsValidSolution(t *testing.T) {
	s, _ := FromFile("testfiles/simple.sudoku")
	s.Solve(SolveOptions{})
	assert.True(t, s.IsValidSolution())
	s.Fields[0].Value = 4
	assert.False(t, s.IsValidSolution())
}

func TestSolveDeduction(t *testing.T) {
	s, _ := FromFile("testfiles/simple.sudoku")
	assert.False(t, s.IsSolved())
	s.Solve(SolveOptions{})
	assert.True(t, s.IsSolved(), "Simple 2x2 sudoku not solved")
}

func TestSolveDeduction3x3(t *testing.T) {
	s, _ := FromFile("testfiles/easy.sudoku")
	assert.False(t, s.IsSolved())
	s.Solve(SolveOptions{})
	assert.True(t, s.IsSolved(), "Easy 3x3 sudoku not solved")
}
