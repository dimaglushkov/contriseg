package animation

import (
	"github.com/dimaglushkov/contriseg/internal"
	"github.com/dimaglushkov/contriseg/internal/image"
	"github.com/dimaglushkov/contriseg/internal/util"
)

func DrawColByColLeft(cal internal.Calendar) []internal.Calendar {
	iterations := []internal.Calendar{cal}
	var orderQueue util.PairQueue
	facingUp := true

	for i := range cal {
		if facingUp {
			for j := len(cal[i]) - 1; j >= 0; j-- {
				if cal[i][j] == 0 {
					cal[i][j] = image.BlueColor
					orderQueue.Push(i, j)
					iterations = append(iterations, cal.Dup())
				}
			}
		} else {
			for j := 0; j < len(cal[i]); j++ {
				if cal[i][j] == 0 {
					cal[i][j] = image.BlueColor
					orderQueue.Push(i, j)
					iterations = append(iterations, cal.Dup())
				}
			}
		}
		facingUp = !facingUp
	}

	for len(orderQueue) > 0 {
		x, y := orderQueue.Pop()
		cal[x][y] = 0
		iterations = append(iterations, cal.Dup())
	}

	return iterations
}
