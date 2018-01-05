package sudoku

import (
	"fmt"
	"strconv"
)

// Field is a sudoku field
type Field struct {
	Index     int
	Value     int
	NonValues *IntSet
	sudoku    *Sudoku
}

// NewField creates a new Field
func NewField(sudoku *Sudoku, index, value int) Field {
	f := Field{
		sudoku:    sudoku,
		Index:     index,
		Value:     value,
		NonValues: NewIntSet(),
	}
	return f
}

// DenyValue denies a value
func (f *Field) DenyValue(value int) {
	f.NonValues.Add(value)
}

// Solve solves this Field
func (f *Field) Solve() bool {
	// if newly solved, return true
	if !f.Solvable() {
		return false
	}
	for i := 1; i <= f.sudoku.MaxValue; i++ {
		if !f.NonValues.Contains(i) {
			f.sudoku.addSolution(f, i)
			break
		}
	}
	return true
}

// IsSolved checks if the Field is solved
func (f Field) IsSolved() bool {
	return f.Value != 0
}

// Solvable checks if the Field can be solved (all other values excluded)
func (f Field) Solvable() bool {
	return len(f.NonValues.set) == f.sudoku.MaxValue-1
}

// PossibleValues returns the list of possible values for this field
func (f Field) PossibleValues() []int {
	result := make([]int, 0)
	for p := 1; p <= f.sudoku.MaxValue; p++ {
		if f.NonValues.Contains(p) {
			continue
		}
		result = append(result, p)
	}
	return result
}

// String returns a human-friendly value
func (f Field) String() string {
	if f.Value == 0 {
		return "."
	}
	fieldLengthString := strconv.Itoa(f.sudoku.FieldLength)
	return fmt.Sprintf("%"+fieldLengthString+"d", f.Value)
}
