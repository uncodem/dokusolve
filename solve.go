package main

import "container/heap"

type CellPosition struct {
	X, Y int
}

func buildCandidates(b Board) [9][9]CandidatePool {
	ret := [9][9]CandidatePool{}

	for y, row := range b {
		for x, cell := range row {
			if cell == 0 {
				ret[y][x] = getCandidates(x, y, b)
			} else {
				ret[y][x] = CandidatePool{Filled: true}
			}
		}
	}

	return ret
}

func buildCandidateHeap(b Board) CandidateHeap {
	ret := make(CandidateHeap, 0, 27)
	heap.Init(&ret)
	candidates := buildCandidates(b)

	for y := range BSIZE {
		for x := range BSIZE {
			pool := candidates[y][x]

			// Allow invalid cells to be pushed to detect the invalid cells easier.
			if !pool.Filled {
				heap.Push(&ret, CandidateHeapEntry{X: x, Y: y, Pool: pool})
			}
		}
	}

	return ret
}

// updatePeers :: int -> int -> int -> CandidateMatrix -> (invalid?, NewSingletons)
func updatePeers(x, y, v int, candidates *[9][9]CandidatePool) (bool, []CellPosition) {
	ret := make([]CellPosition, 0, 27)
	sqx := (x / 3) * 3
	sqy := (y / 3) * 3

	// Row

	for col := range BSIZE {
		if col == x {
			continue
		}
		cell := &candidates[y][col]
		if !cell.Filled {
			cell.Remove(v)
			if cell.Count == 0 {
				return true, nil
			} else if cell.Count == 1 {
				ret = append(ret, CellPosition{X: col, Y: y})
			}
		}
	}

	// Column
	for row := range BSIZE {
		if row == y {
			continue
		}
		cell := &candidates[row][x]
		if !cell.Filled {
			cell.Remove(v)
			if cell.Count == 0 {
				return true, nil
			} else if cell.Count == 1 {
				ret = append(ret, CellPosition{X: x, Y: row})
			}
		}
	}

	// Square
	for sy := range 3 {
		for sx := range 3 {
			if sqy+sy == y && sqx+sx == x {
				continue
			}

			cell := &candidates[sqy+sy][sqx+sx]
			if !cell.Filled {
				removed := cell.Remove(v)
				if removed && cell.Count == 0 {
					return true, nil
				} else if removed && cell.Count == 1 {
					ret = append(ret, CellPosition{X: sqx + sx, Y: sqy + sy})
				}
			}
		}
	}
	return false, ret
}

// singletonSweep :: Board -> (Board' * Solvable?)
func singletonSweep(b Board) (Board, bool) {
	cloned := b
	candidates := buildCandidates(b)

	for y := range BSIZE {
		for x := range BSIZE {
			peers := make([]CellPosition, 1, 27)
			peers[0] = CellPosition{X: x, Y: y}
			for i := 0; i < len(peers); i++ {
				px, py := peers[i].X, peers[i].Y
				cell := &candidates[peers[i].Y][peers[i].X]
				if v := cell.Singleton(); v != -1 { // Singleton detected, value is in v
					cloned[py][px] = v
					cell.Filled = true
					invalidated, new_peers := updatePeers(px, py, v, &candidates)
					if invalidated {
						return b, false
					}
					peers = append(peers, new_peers...)
				}
			}
		}
	}

	return cloned, true
}
