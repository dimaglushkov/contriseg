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
	for _, c := range ColorScheme {
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
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			frame.Set(i, j, ColorScheme[BackgroundColor])
		}
	}
	for i := range cal {
		for j := range cal[i] {
			drawDay(frame, &ColorScheme[cal[i][j]], i, j)
		}
	}
	return frame
}

func GetFrames(cal internal.Calendar, iterator AnimationFunc) ([]*image.Paletted, error) {
	numOfWeeks := len(cal)
	gapW, gapH = (w-numOfWeeks*dayW)/(numOfWeeks-1), (h-7*dayH)/6

	var frames []*image.Paletted

	pal := generatePalette()
	for _, c := range iterator(cal) {
		frames = append(frames, generateFrameFromCal(pal, c))
	}

	return frames, nil
}
