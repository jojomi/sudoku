package main

import (
	"fmt"

	"github.com/jojomi/sudoku"
)

func main() {
	/*s := sudoku.New(2)
	s.Init([]int{
		1, 0, 2, 3,
		0, 0, 4, 0,
		3, 2, 0, 0,
		0, 0, 0, 0,
		//
		//	1, 4, 2, 3, // 0, 1, 2, 3
		//	2, 3, 4, 1, // 4, 5, 6, 7
		//	3, 2, 1, 4, // 8, 9, 10, 11
		//	4, 1, 3, 2, // 12, 13, 14, 15
		//
	})
	fmt.Println(s)
	//s.Explain()
	s.Solve()
	//s.Solve()
	fmt.Println(s)*/

	/*s3 := sudoku.New(3)
	s3.Init([]int{
		9, 6, 0, 5, 0, 0, 1, 0, 4,
		0, 0, 0, 4, 0, 0, 5, 6, 0,
		5, 4, 2, 0, 0, 1, 0, 0, 0,
		3, 0, 0, 0, 0, 0, 0, 1, 6,
		0, 0, 1, 9, 0, 7, 0, 0, 5,
		0, 0, 0, 1, 0, 0, 0, 0, 8,
		0, 0, 4, 0, 0, 0, 0, 0, 1,
		0, 0, 0, 0, 2, 3, 4, 5, 7,
		0, 5, 0, 0, 1, 4, 6, 0, 0,
	})
	fmt.Println(s3)
	for !s3.IsSolved() {
		s3.Solve()
	}
	fmt.Println(s3)*/

	s4 := sudoku.FromFile("test/neuner.sudoku")
	fmt.Println(s4)
	s4.Solve()
	fmt.Println(s4)
}
