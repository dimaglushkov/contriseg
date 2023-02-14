package util

type pair struct {
	x, y int
}

type PairQueue []pair

func (q *PairQueue) Push(x, y int) {
	*q = append(*q, pair{x, y})
}
func (q *PairQueue) Pop() (int, int) {
	x := (*q)[0]
	*q = (*q)[1:]
	return x.x, x.y
}
