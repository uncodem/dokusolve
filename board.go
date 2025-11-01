package main

import "fmt"

const (
	BSIZE = 9
)

type Board [BSIZE][BSIZE]int

func readBoard() Board {
	ret := Board{}
	for y := range BSIZE {
		for x := range BSIZE {
			fmt.Scanf("%d", &ret[y][x])
		}
	}
	return ret
}

func printBoard(b Board) {
	hline := "+-------+-------+-------+"
	for y := range BSIZE {
		if y%3 == 0 {
			fmt.Println(hline)
		}
		for x := range BSIZE {
			if x%3 == 0 {
				fmt.Print("| ")
			}
			if b[y][x] == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", b[y][x])
			}
		}
		fmt.Println("| ")
	}
	fmt.Println(hline)
}

func getCandidates(x, y int, b Board) CandidatePool {
	ret := CandidatePool{
		Count:     0,
		Pool:      [9]bool{true, true, true, true, true, true, true, true, true},
		Filled:    false,
		PoolSlice: nil,
	}

	sqx := (x / 3) * 3
	sqy := (y / 3) * 3

	// Column
	for row := range BSIZE {
		cell := b[row][x]
		if cell != 0 {
			ret.Pool[cell-1] = false
		}
	}

	// Row
	for col := range BSIZE {
		cell := b[y][col]
		if cell != 0 {
			ret.Pool[cell-1] = false
		}
	}

	for sy := range 3 {
		for sx := range 3 {
			cell := b[sqy+sy][sqx+sx]
			if cell != 0 {
				ret.Pool[cell-1] = false
			}
		}
	}

	for _, candidate := range ret.Pool {
		if candidate {
			ret.Count++
		}
	}
	return ret
}

func validBoard(b Board) bool {
	candidates := buildCandidates(b)
	for y := range BSIZE {
		for x := range BSIZE {
			cell := b[y][x]

			if cell == 0 {
				if candidates[y][x].Count == 0 {
					return false
				}
				continue
			}

			// Row
			for col := range BSIZE {
				if col != x && b[y][col] == cell {
					return false
				}
			}

			// Column
			for row := range BSIZE {
				if row != y && b[row][x] == cell {
					return false
				}
			}

			// 3×3 square
			sqx := (x / 3) * 3
			sqy := (y / 3) * 3
			for sy := range 3 {
				for sx := range 3 {
					yy, xx := sqy+sy, sqx+sx
					if yy == y && xx == x {
						continue
					}
					if b[yy][xx] == cell {
						return false
					}
				}
			}
		}
	}

	return true
}

func solved(b Board) bool {
	for _, row := range b {
		for _, column := range row {
			if column == 0 {
				return false
			}
		}
	}
	return true
}
