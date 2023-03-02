package animation

import (
	"github.com/dimaglushkov/contriseg/internal"
	"github.com/dimaglushkov/contriseg/internal/image"
	"github.com/dimaglushkov/contriseg/internal/util"
)

func DrawBFS(cal internal.Calendar) []internal.Calendar {
	iterations := []internal.Calendar{cal}
	numOfWeeks := len(cal)
	var queue util.PairQueue
	var orderQueue util.PairQueue
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < numOfWeeks; i++ {
		for j := 0; j < len(cal[i]); j++ {
			if cal[i][j] == 0 {
				queue.Push(i, j)
				orderQueue.Push(i, j)
				for len(queue) > 0 {
					x, y := queue.Pop()
					if cal[x][y] != 0 {
						continue
					}
					cal[x][y] = image.BlueColor
					iterations = append(iterations, cal.Dup())

					orderQueue.Push(x, y)
					for _, d := range directions {
						if nx, ny := x+d[0], y+d[1]; nx >= 0 && nx < numOfWeeks && ny >= 0 && ny < len(cal[nx]) && cal[nx][ny] == 0 {
							queue.Push(nx, ny)
						}
					}
				}
			}
		}
	}
	for len(orderQueue) > 0 {
		x, y := orderQueue.Pop()
		cal[x][y] = 0
		iterations = append(iterations, cal.Dup())
	}

	return iterations
}
