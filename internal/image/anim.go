package image

import (
	"fmt"
	"github.com/dimaglushkov/contriseg/internal"
	"github.com/dimaglushkov/contriseg/internal/util"
	"strings"
)

type AnimationIterator func(cal internal.Calendar) []internal.Calendar

var animationsMap = map[string]AnimationIterator{
	"bfs":  CalendarBFSIterations,
	"move": CalendarMoveColLeftIterations,
	"cbc":  CalendarColByColIterations,
}

func GetAvailableAnimations() []string {
	availableAnimations := make([]string, 0, len(animationsMap))
	for k := range animationsMap {
		availableAnimations = append(availableAnimations, k)
	}
	return availableAnimations
}

func GetAnimationIterator(animation string) (AnimationIterator, error) {
	var iter AnimationIterator
	var ok bool

	if iter, ok = animationsMap[strings.ToLower(animation)]; !ok {
		return nil, fmt.Errorf("unknown framing animation: %s, available animations are: %s (case insesnsetive)", animation, strings.Join(GetAvailableAnimations(), ", "))
	}

	return iter, nil
}

func CalendarBFSIterations(cal internal.Calendar) []internal.Calendar {
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
					cal[x][y] = BlueColor
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

func CalendarColByColIterations(cal internal.Calendar) []internal.Calendar {
	iterations := []internal.Calendar{cal}
	var orderQueue util.PairQueue
	facingUp := true

	for i := range cal {
		if facingUp {
			for j := len(cal[i]) - 1; j >= 0; j-- {
				if cal[i][j] == 0 {
					cal[i][j] = BlueColor
					orderQueue.Push(i, j)
					iterations = append(iterations, cal.Dup())
				}
			}
		} else {
			for j := 0; j < len(cal[i]); j++ {
				if cal[i][j] == 0 {
					cal[i][j] = BlueColor
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
