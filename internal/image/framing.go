package image

import (
	"github.com/dimaglushkov/contriseg/internal"
	"github.com/dimaglushkov/contriseg/internal/util"
	"image"
	"image/color"
)

const (
	w       = 800
	h       = 100
	dayW    = 11
	dayH    = 11
	offsetW = 5
	offsetH = 2
)

var gapW, gapH = 3, 4

func generatePalette() []color.Color {
	var palette []color.Color
	palette = append(palette, BackgroundColor)
	for _, c := range GithubContribScheme {
		palette = append(palette, c)
	}
	for _, c := range GradientScheme {
		palette = append(palette, c)
	}
	return palette
}

func drawDay(frame *image.Paletted, c *color.RGBA, weekNum, dayNum int) {
	sx, sy := (dayW+gapW)*weekNum+offsetW, (dayH+gapH)*dayNum+offsetH
	for i := 0; i < dayW; i++ {
		for j := 0; j < dayH; j++ {
			frame.Set(sx+i, sy+j, c)
		}
	}
}

func generateFrameFromCal(palette []color.Color, cal internal.Calendar) *image.Paletted {
	frame := image.NewPaletted(image.Rect(0, 0, w, h), palette)
	for i := range cal {
		for j := range cal[i] {
			drawDay(frame, &BackgroundColor, i, j)
		}
	}

	for i := range cal {
		for j := range cal[i] {
			if cal[i][j] == -1 {
				drawDay(frame, &DarkBlueColor, i, j)
			} else {
				drawDay(frame, &GithubContribScheme[cal[i][j]], i, j)
			}
		}
	}
	return frame
}

func BFSFrames(cal internal.Calendar) ([]*image.Paletted, error) {
	numOfWeeks := len(cal)
	gapW, gapH = (w-numOfWeeks*dayW)/(numOfWeeks-1), (h-7*dayH)/6

	var frames []*image.Paletted
	var queue util.PairQueue
	var orderQueue util.PairQueue

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	pal := generatePalette()
	frames = append(frames, generateFrameFromCal(pal, cal))

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
					cal[x][y] = -1
					frames = append(frames, generateFrameFromCal(pal, cal))
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
		frames = append(frames, generateFrameFromCal(pal, cal))
	}

	return frames, nil
}
