package main

import "container/heap"

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
	ret := make(CandidateHeap, 0)
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

// updatePeers :: int -> int -> int -> CandidateMatrix -> invalid?
func updatePeers(x, y, v int, candidates *[9][9]CandidatePool) bool {
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
				return true
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
				return true
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
				cell.Remove(v)
				if cell.Count == 0 {
					return true
				}
			}
		}
	}
	return false
}

// singletonSweep :: Board -> (Board' * Solvable?)
func singletonSweep(b Board) (Board, bool) {
	cloned := b
	candidates := buildCandidates(b)

	for {
		changed := false
		for y := range BSIZE {
			for x := range BSIZE {
				cell := &candidates[y][x]
				if v := cell.Singleton(); v != -1 { // Singleton detected, value is in v
					cloned[y][x] = v
					cell.Filled = true
					changed = true
					if updatePeers(x, y, v, &candidates) {
						return b, false
					}
				}
			}
		}
		if !changed {
			break
		}
	}

	return cloned, true
}
