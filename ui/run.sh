#!/bin/sh

set -e

go build -o sudoku
./sudoku "$@"
