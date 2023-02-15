package internal

// Calendar represents contributions calendar by precalculated day quartiles
type Calendar [][]int8

func (c Calendar) Dup() Calendar {
	dup := make(Calendar, len(c))
	for i := range c {
		dup[i] = make([]int8, len(c[i]))
		copy(dup[i], c[i])
	}
	return dup
}
