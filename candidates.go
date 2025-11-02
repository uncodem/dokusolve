package main

type CandidatePool struct {
	Filled    bool
	Count     int
	Pool      [9]bool
	PoolSlice []int
}

type CandidateHeapEntry struct {
	X, Y int
	Pool CandidatePool
}

type CandidateHeap []CandidateHeapEntry

func (h CandidateHeap) Len() int           { return len(h) }
func (h CandidateHeap) Less(i, j int) bool { return h[i].Pool.Count < h[j].Pool.Count }
func (h CandidateHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CandidateHeap) Push(x any) {
	*h = append(*h, x.(CandidateHeapEntry))
}

func (h *CandidateHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (cp *CandidatePool) CalcSlice() {
	cp.PoolSlice = make([]int, 0)
	for i, candidate := range cp.Pool {
		if candidate {
			cp.PoolSlice = append(cp.PoolSlice, i+1)
		}
	}
}

func (cp *CandidatePool) Remove(num int) bool {
	if cp.Pool[num-1] {
		cp.Pool[num-1] = false
		cp.Count -= 1
		if cp.PoolSlice != nil {
			cp.CalcSlice()
		}
		return true
	}
	return false
}

func (cp CandidatePool) Singleton() int {
	if cp.Filled {
		return -1
	}

	ret := -1
	if cp.Count == 1 {
		for i, n := range cp.Pool {
			if n {
				ret = i + 1
				break
			}
		}
	}
	return ret
}
