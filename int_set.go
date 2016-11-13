package sudoku

import "fmt"

// IntSet is a set of integers
type IntSet struct {
	set map[int]bool
}

// NewIntSet makes a new IntSet instance
func NewIntSet() *IntSet {
	i := &IntSet{
		set: make(map[int]bool),
	}
	return i
}

// Add an entry
func (set *IntSet) Add(i int) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found // false if it existed already
}

// Contains tells if the integer is already in the set
func (set *IntSet) Contains(i int) bool {
	_, found := set.set[i]
	return found
}

// String returns a string representation
func (set *IntSet) String() string {
	return fmt.Sprintln(set.set)
}
