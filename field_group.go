package sudoku

import "fmt"

// FieldGroup is a set of Fields (Row/Col/Block)
type FieldGroup struct {
	sudoku *Sudoku
	Fields []*Field
	Name   string
}

// NewFieldGroup makes a new fieldgroup of specified size
func NewFieldGroup(sudoku *Sudoku, len int, name string) FieldGroup {
	fg := FieldGroup{
		sudoku: sudoku,
		Fields: make([]*Field, len),
		Name:   name,
	}
	return fg
}

// IsSolved checks if this FG is solved
func (f FieldGroup) IsSolved() bool {
	for _, field := range f.Fields {
		if !field.IsSolved() {
			return false
		}
	}
	return true
}

// Solve solves the group if possible (one step!)
func (f FieldGroup) Solve() SolvingResult {
	// loop possible values
valueLoop:
	for val := 1; val <= f.sudoku.MaxValue; val++ {
		deducedField := -1
		// loop fields, count options
		for _, field := range f.Fields {
			if field.Value == val {
				continue valueLoop
			}
			if field.IsSolved() {
				continue
			}
			if !field.NonValues.Contains(val) {
				if deducedField != -1 {
					continue valueLoop
				} else {
					deducedField = field.Index
				}
			}
		}
		// if exactly one option, solve!
		if deducedField != -1 {
			f.sudoku.addSolutionByIndex(deducedField, val)
			return SolvingResult{
				FoundNew: true,
				Message:  fmt.Sprintf("Deduced by checking %s: Field %d must be of value %d", f.Name, deducedField, val),
			}
		}
	}
	return SolvingResult{}
}

// String returns a string representation of this FieldGroup
func (f FieldGroup) String() string {
	result := ""
	for _, field := range f.Fields {
		result += field.String()
	}
	return result
}
