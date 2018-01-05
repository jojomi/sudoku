package sudoku

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
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

type SolveOptions struct {
	PrintSteps bool
	DeduceOnly bool
	DontDeduce bool
}

// New returns a new sudoku puzzle
func New(size int) *Sudoku {
	s := &Sudoku{
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

// FromFile loads a sudoku definition from a file
func FromFile(filename string) (*Sudoku, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return FromReader(strings.NewReader(string(data)))
}

func FromReader(input io.Reader) (*Sudoku, error) {
	data, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}
	// clean data
	cleanString := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, string(data))

	mathSize := math.Sqrt(math.Sqrt(float64(len(cleanString))))
	size := int(mathSize)
	// TODO check size

	initData := make([]int, int(math.Pow(float64(size), 4)))
	for i, c := range cleanString {
		initData[i], err = strconv.Atoi(string(c))
		if err != nil {
			// this allows for all non-whitespace chars to signal empty fields
			initData[i] = 0
		}
	}

	s := New(size)
	s.Init(initData)
	return s, nil
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

// Reason
func (s Sudoku) Reason() Sudoku {
	for _, f := range s.Fields {
		if f.IsSolved() {
			s.addSolution(f, f.Value)
		}
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

func (s Sudoku) SolvedFieldCount() int {
	result := 0
	for _, f := range s.Fields {
		if f.IsSolved() {
			result++
		}
	}
	return result
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

// IsValidSolution checks a sudoku for validity
func (s Sudoku) IsValidSolution() bool {
	var set *IntSet

	for _, row := range s.rows {
		set = NewIntSet()
		for _, f := range row.Fields {
			if f.IsSolved() {
				set.Add(f.Value)
			}
		}
		if len(set.Values()) < s.MaxValue {
			return false
		}
	}

	for _, col := range s.cols {
		set = NewIntSet()
		for _, f := range col.Fields {
			if f.IsSolved() {
				set.Add(f.Value)
			}
		}
		if len(set.Values()) < s.MaxValue {
			return false
		}
	}

	for _, block := range s.blocks {
		set = NewIntSet()
		for _, f := range block.Fields {
			if f.IsSolved() {
				set.Add(f.Value)
			}
		}
		if len(set.Values()) < s.MaxValue {
			return false
		}
	}

	return true
}

// Solve solves
func (s Sudoku) Solve(opts SolveOptions) Sudoku {
	res := SolvingResult{
		FoundNew: true,
	}

	// what do we already know?
	s.Reason()

	// Phase 1: Deduction
	if !opts.DontDeduce {
		for !s.IsSolved() && res.FoundNew {
			res = s.SolveStep(opts)
			if opts.PrintSteps {
				fmt.Println(res)
				fmt.Println(s)
			}
		}
	}
	// Phase 2: Backtracking
	if !s.IsSolved() {
		s.SolveBrute(opts)
	}
	return s
}

type SolvingResult struct {
	FoundNew bool
	Message  string
}

func (s SolvingResult) String() string {
	return s.Message
}

// SolveStep solves one step
func (s Sudoku) SolveStep(opts SolveOptions) SolvingResult {
	var res SolvingResult

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
				res = SolvingResult{
					FoundNew: true,
					Message:  fmt.Sprintf("Field %d could be deduced because there was only one more possible value which is %d", f.Index, f.Value),
				}
				return res
			}
		}
	}

	// loop all three dimensions
	//   loop all valid numbers
	//     loop all fields for checking
	//       solve and go back to start
	for _, row := range s.rows {
		res = row.Solve()
		if res.FoundNew {
			return res
		}
	}
	for _, col := range s.cols {
		res = col.Solve()
		if res.FoundNew {
			return res
		}
	}
	for _, block := range s.blocks {
		res = block.Solve()
		if res.FoundNew {
			return res
		}
	}
	return res
}

func (s Sudoku) UnsolvedFields() []*Field {
	result := make([]*Field, 0)
	for _, f := range s.Fields {
		if !f.IsSolved() {
			result = append(result, f)
		}
	}
	return result
}

func (s Sudoku) CanPut(field *Field, value int) bool {
	for _, f := range s.GetRow(field).Fields {
		if f.Value == value {
			return false
		}
	}
	for _, f := range s.GetCol(field).Fields {
		if f.Value == value {
			return false
		}
	}
	for _, f := range s.GetBlock(field).Fields {
		if f.Value == value {
			return false
		}
	}
	return true
}

// SolveBrute brute-forces a sudoku
func (s Sudoku) SolveBrute(options SolveOptions) bool {
	if options.PrintSteps {
		fmt.Println("I need brute force.")
	}
	// missing fields?
	fields := s.UnsolvedFields()

	// solved?
	return s.solveBruteStep(options, fields)
}

func (s Sudoku) solveBruteStep(options SolveOptions, fields []*Field) bool {
	if len(fields) == 0 {
		return s.IsValidSolution()
	}
	f := fields[0]

	for _, v := range f.PossibleValues() {
		if options.PrintSteps {
			fmt.Printf("trying %d at field %d\n", v, f.Index)
		}
		if !s.CanPut(f, v) {
			continue
		}
		f.Value = v
		if s.solveBruteStep(options, fields[1:]) {
			return true
		}
	}
	f.Value = 0

	return false
}
