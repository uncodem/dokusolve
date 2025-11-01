package main

// phaseTwo :: Board -> (Board' * Solvable?)

func phaseTwo(b Board) (Board, bool) {
	clone := b
	stack := make([]Board, 0)
	stack = append(stack, clone)

	for len(stack) != 0 {
		board := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if solved(board) {
			return board, true
		}

		candidates := buildCandidateHeap(board)
		if len(candidates) == 0 {
			continue
		}

		if candidates[0].Pool.Count == 0 {
			continue
		}

		cell := candidates[0]
		cell.Pool.CalcSlice()

		for _, candidate := range cell.Pool.PoolSlice {
			clone := board
			clone[cell.Y][cell.X] = candidate
			clone, solvable := singletonSweep(clone)
			if solvable {
				stack = append(stack, clone)
			}
		}

	}

	return b, false
}
