package sudoku

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Optimization ideas:
// - keep track of solved/unsolved fields (shorter loops, quicker check for solved state)
// - auto solve on adding n-1th element in deny list
// - dirty state for row/col/block?

// Sudoku is a sudoku puzzle
type Sudoku struct {
	Size        int
	MaxValue    int
	FieldLength int
	Fields      []*Field
	cols        []FieldGroup
	rows        []FieldGroup
	blocks      []FieldGroup
}

// New returns a new sudoku puzzle
func New(size int) Sudoku {
	s := Sudoku{
		Size:     size,
		MaxValue: size * size,
	}
	fieldLength := len(strconv.Itoa(s.MaxValue))
	s.FieldLength = fieldLength
	// init fields
	fieldCount := size * size * size * size

	lineSize := s.MaxValue
	s.rows = make([]FieldGroup, lineSize)
	s.cols = make([]FieldGroup, lineSize)
	s.blocks = make([]FieldGroup, lineSize)

	s.Fields = make([]*Field, fieldCount)
	for row := 0; row < lineSize; row++ {
		for col := 0; col < lineSize; col++ {
			index := row*lineSize + col
			field := NewField(s, index, 0)
			s.Fields[index] = &field
			if col == 0 {
				s.rows[row] = NewFieldGroup(s, lineSize, fmt.Sprintf("row %d", row))
			}
			s.rows[row].Fields[col] = &field
			if row == 0 {
				s.cols[col] = NewFieldGroup(s, lineSize, fmt.Sprintf("col %d", col))
			}
			s.cols[col].Fields[row] = &field

			blockIndex := int(math.Floor(float64(row)/float64(s.Size)))*s.Size + int(math.Floor(float64(col)/float64(s.Size)))
			innerBlockRow := int(math.Mod(float64(row), float64(s.Size)))
			innerBlockCol := int(math.Mod(float64(col), float64(s.Size)))
			isBlockStart := innerBlockRow == 0 && innerBlockCol == 0
			if isBlockStart {
				s.blocks[blockIndex] = NewFieldGroup(s, lineSize, fmt.Sprintf("block %d", blockIndex))
			}
			s.blocks[blockIndex].Fields[innerBlockRow*s.Size+innerBlockCol] = &field
		}
	}
	return s
}

// IsSolved checks if this sudoku is solved
func (s Sudoku) IsSolved() bool {
	for _, f := range s.Fields {
		if !f.IsSolved() {
			return false
		}
	}
	return true
}

// String returns the state of the sudoku
func (s Sudoku) String() string {
	result := ""
	lineSize := s.MaxValue
	drawBorder := true

	// top border
	if drawBorder {
		result += "+" + strings.Repeat(strings.Repeat("-", s.Size)+"+", s.Size) + "\n"
	}
	for row := 0; row < lineSize; row++ {
		if drawBorder {
			result += "|"
		}
		for col := 0; col < lineSize; col++ {
			index := row*lineSize + col
			field := s.Fields[index]
			result += field.String()
			// end of block?
			if drawBorder && col%s.Size == s.Size-1 {
				result += "|"
			}
		}
		result += "\n"
		if drawBorder && row%s.Size == s.Size-1 {
			result += "+" + strings.Repeat(strings.Repeat("-", s.Size)+"+", s.Size) + "\n"
		}
	}

	return result
}

// Init inits
func (s Sudoku) Init(input []int) Sudoku {
	for i, val := range input {
		s.Fields[i].Value = val
	}
	return s
}

// GetRow returns all fields of the same row as given field
func (s Sudoku) GetRow(f *Field) FieldGroup {
	row := int(math.Floor(float64(f.Index) / float64(s.MaxValue)))
	return s.rows[row]
}

// GetCol returns all fields of the same column as given field
func (s Sudoku) GetCol(f *Field) FieldGroup {
	col := int(math.Mod(float64(f.Index), float64(s.MaxValue)))
	return s.cols[col]
}

// GetBlock gets all block Fields for a given field
func (s Sudoku) GetBlock(f *Field) FieldGroup {
	row := int(math.Floor(float64(f.Index) / float64(s.MaxValue)))
	col := int(math.Mod(float64(f.Index), float64(s.MaxValue)))
	blockIndex := int(math.Floor(float64(row)/float64(s.Size)))*s.Size + int(math.Floor(float64(col)/float64(s.Size)))

	return s.blocks[blockIndex]
}

func (s Sudoku) addSolutionByIndex(index, value int) {
	s.addSolution(s.Fields[index], value)
}

func (s Sudoku) addSolution(f *Field, value int) {
	f.Value = value

	rowFields := s.GetRow(f)
	for _, rf := range rowFields.Fields {
		if rf.IsSolved() {
			continue
		}
		//fmt.Printf("Field %d can't be of value %d (row check)\n", rf.Index, f.Value)
		rf.DenyValue(f.Value)
	}

	colFields := s.GetCol(f)
	for _, cf := range colFields.Fields {
		if cf.IsSolved() {
			continue
		}
		//fmt.Printf("Field %d can't be of value %d (col check)\n", cf.Index, f.Value)
		cf.DenyValue(f.Value)
	}

	blockFields := s.GetBlock(f)
	for _, bf := range blockFields.Fields {
		if bf.IsSolved() {
			continue
		}
		//fmt.Printf("Field %d can't be of value %d (block check)\n", bf.Index, f.Value)
		bf.DenyValue(f.Value)
	}
}

// Solve solves
func (s Sudoku) Solve() Sudoku {
	// loop all solved fields at begin
	//   create falsity information along all three dimensions
	for _, f := range s.Fields {
		if f.IsSolved() {
			s.addSolution(f, f.Value)
		}
	}

	for _, f := range s.Fields {
		// not solved
		if !f.IsSolved() {
			if f.Solve() {
				fmt.Printf("Field %d could be deduced because there was only one more possible value which is %d\n", f.Index, f.Value)
			}
		}
	}

	// loop all three dimensions
	//   loop all valid numbers
	//     loop all fields for checking
	//       solve and go back to start
	for _, row := range s.rows {
		row.Solve()
	}
	for _, col := range s.cols {
		col.Solve()
	}
	for _, block := range s.blocks {
		block.Solve()
	}
	return s
}
