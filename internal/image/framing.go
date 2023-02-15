package image

import (
	"github.com/dimaglushkov/contriseg/internal"
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

var (
	gapW = 3
	gapH = 4
)

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

func GetFrames(cal internal.Calendar, iterator AnimationIterator) ([]*image.Paletted, error) {
	numOfWeeks := len(cal)
	gapW, gapH = (w-numOfWeeks*dayW)/(numOfWeeks-1), (h-7*dayH)/6

	var frames []*image.Paletted

	pal := generatePalette()
	for _, c := range iterator(cal, true) {
		frames = append(frames, generateFrameFromCal(pal, c))
	}

	return frames, nil
}
