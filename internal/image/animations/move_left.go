package animations

import "github.com/dimaglushkov/contriseg/internal"

func CalendarMoveColLeftIterations(cal internal.Calendar) []internal.Calendar {
	iterations := []internal.Calendar{cal}
	n := len(cal)
	tmp, last := make([]int8, 7), make([]int8, 7)
	for range cal {
		copy(last, tmp)
		copy(tmp, cal[0])
		for j := 0; j < n-2; j++ {
			copy(cal[j], cal[j+1])
		}
		copy(cal[n-2], last)
		copy(cal[n-1], tmp[:len(cal[n-1])])
		iterations = append(iterations, cal.Dup())
	}

	return iterations
}
