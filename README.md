# dokusolve
A two-phase sudoku solver written in Go.   

The first phase fills cells that could only have one logical solution. If the puzzle is not finished, it proceeds to the second phase. The second phase utilizes a stack of board states to avoid recursion and a min-heap of cells and their candidate pools to search for the solution without exploring every single possibility. 
## Usage
This program reads a puzzle from STDIN and solves it within 4-10ms. 
```
> cat puzzle.txt
0 0 8 2 0 0 0 4 0
0 0 7 0 0 0 8 6 0
2 0 0 0 0 8 0 0 0
6 0 0 0 0 1 0 0 0
0 9 0 0 4 0 0 5 0
0 0 0 3 0 0 0 0 7
0 0 0 8 0 0 0 0 5
0 4 5 0 0 0 7 0 0
0 3 0 0 0 9 1 0 0

> go run . < puzzle.txt
Unsolved:
+-------+-------+-------+
| . . 8 | 2 . . | . 4 . |
| . . 7 | . . . | 8 6 . |
| 2 . . | . . 8 | . . . |
+-------+-------+-------+
| 6 . . | . . 1 | . . . |
| . 9 . | . 4 . | . 5 . |
| . . . | 3 . . | . . 7 |
+-------+-------+-------+
| . . . | 8 . . | . . 5 |
| . 4 5 | . . . | 7 . . |
| . 3 . | . . 9 | 1 . . |
+-------+-------+-------+
Solved:
+-------+-------+-------+
| 3 5 8 | 2 7 6 | 9 4 1 |
| 4 1 7 | 9 3 5 | 8 6 2 |
| 2 6 9 | 4 1 8 | 5 7 3 |
+-------+-------+-------+
| 6 7 2 | 5 8 1 | 4 3 9 |
| 1 9 3 | 6 4 7 | 2 5 8 |
| 5 8 4 | 3 9 2 | 6 1 7 |
+-------+-------+-------+
| 7 2 1 | 8 6 4 | 3 9 5 |
| 9 4 5 | 1 2 3 | 7 8 6 |
| 8 3 6 | 7 5 9 | 1 2 4 |
+-------+-------+-------+
```
