# sudoku
Solving sudoku puzzles in Golang


# Execute

    cd ui
    go build main.go && time sudo ./main


# Example output

```
96.5..1.4
...4..56.
542..1...
3......16
..19.7..5
...1....8
..4.....1
....23457
.5..146..

Deduced by checking row 3: Field 31 must be of value 4
Deduced by checking col 3: Field 30 must be of value 2
Deduced by checking col 3: Field 21 must be of value 3
Deduced by checking block 1: Field 22 must be of value 6
Deduced by checking block 4: Field 50 must be of value 6
Field 26 could be deduced because there was only one more possible value which is 9
Deduced by checking row 4: Field 36 must be of value 6
Deduced by checking row 6: Field 57 must be of value 6
Deduced by checking row 7: Field 65 must be of value 6
Deduced by checking row 7: Field 64 must be of value 9
Deduced by checking row 8: Field 79 must be of value 9
Deduced by checking col 0: Field 45 must be of value 4
Deduced by checking col 1: Field 10 must be of value 1
Deduced by checking col 1: Field 55 must be of value 3
Deduced by checking col 3: Field 75 must be of value 7
Deduced by checking col 3: Field 66 must be of value 8
Deduced by checking col 7: Field 43 must be of value 4
Deduced by checking block 6: Field 63 must be of value 1
Deduced by checking block 6: Field 54 must be of value 7
Deduced by checking block 8: Field 80 must be of value 3
Field 9 could be deduced because there was only one more possible value which is 8
Field 17 could be deduced because there was only one more possible value which is 2
Field 72 could be deduced because there was only one more possible value which is 2
Field 74 could be deduced because there was only one more possible value which is 8
Deduced by checking row 0: Field 5 must be of value 2
Deduced by checking row 1: Field 11 must be of value 3
Deduced by checking row 1: Field 13 must be of value 7
Deduced by checking row 1: Field 14 must be of value 9
Deduced by checking row 6: Field 58 must be of value 9
Deduced by checking col 4: Field 49 must be of value 5
Deduced by checking col 5: Field 59 must be of value 5
Deduced by checking col 5: Field 32 must be of value 8
Deduced by checking block 0: Field 2 must be of value 7
Deduced by checking block 1: Field 4 must be of value 8
Deduced by checking block 2: Field 7 must be of value 3
Deduced by checking block 3: Field 29 must be of value 5
Deduced by checking block 3: Field 37 must be of value 8
Deduced by checking block 3: Field 47 must be of value 9
Deduced by checking block 4: Field 40 must be of value 3
Deduced by checking block 5: Field 51 must be of value 3
Deduced by checking block 5: Field 33 must be of value 9
Field 28 could be deduced because there was only one more possible value which is 7
Field 42 could be deduced because there was only one more possible value which is 2
Field 46 could be deduced because there was only one more possible value which is 2
Field 52 could be deduced because there was only one more possible value which is 7
Field 60 could be deduced because there was only one more possible value which is 8
Field 61 could be deduced because there was only one more possible value which is 2
Deduced by checking row 2: Field 24 must be of value 7
Deduced by checking row 2: Field 25 must be of value 8
967582134
813479562
542361789
375248916
681937245
429156378
734695821
196823457
258714693

real    0m0,034s
user    0m0,007s
sys     0m0,009s
```


## State of the Union

Basically this code can solve any sudoku out there that is solvable. If there is multiple solutions to a puzzle, this program will return one of them only.


### How to Improve?

The tool could be more clever in **calculating** solutions rather then brute forcing. There is a good list of tricks listed [here](https://www.sudokuoftheday.com/techniques/). Currently techniques 1, 2, and 12 are implemented. If there would be more tricks in the codebase, the solver could explain his way to solving better (`-p` flag).

If you are interested in teaching your computer some of those tricks, please do so and file a Pull Request, so everyone can profit.