package main

import (
	"fmt"
	"os"

	"github.com/jojomi/sudoku"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "sudoku filename",
		Run: cmdSolve,
	}

	rootCmd.Execute()
}

func cmdSolve(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Need input filename. Aborting.")
		os.Exit(1)
	}

	s := sudoku.FromFile(args[0])
	fmt.Println("parsed sudoku from input:")
	fmt.Println(s)
	s.Solve()
	fmt.Println("solution:")
	fmt.Println(s)
}
