package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jojomi/sudoku"
	"github.com/spf13/cobra"
)

var (
	solveOptionsPrintSteps bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "sudoku filename",
		Run: cmdSolve,
	}
	rootCmd.PersistentFlags().BoolVarP(&solveOptionsPrintSteps, "print-steps", "p", false, "print steps while solving sudoku")

	rootCmd.Execute()
}

func cmdSolve(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Need input filename. Aborting.")
		os.Exit(1)
	}

	s, err := sudoku.FromFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	opts := sudoku.SolveOptions{
		PrintSteps: solveOptionsPrintSteps,
	}

	fmt.Println("parsed sudoku from input:")
	fmt.Println(s)
	s.Solve(opts)
	fmt.Println("solution:")
	fmt.Println(s)
}
