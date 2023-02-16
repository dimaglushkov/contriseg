package image

import (
	"image/color"
)

const (
	GithubNoContribColor = iota
	GithubFirstQuadrantColor
	GithubSecondQuadrantColor
	GithubThirdQuadrantColor
	GithubFourthQuadrantColor
	BackgroundColor
	DarkBlueColor
	BlueColor
)

var ColorScheme = []color.RGBA{
	{24, 28, 36, 1},
	{14, 68, 41, 1},
	{0, 109, 50, 1},
	{38, 166, 65, 1},
	{57, 211, 83, 1},
	{13, 17, 23, 1},
	{16, 33, 97, 1},
	{10, 30, 150, 1},
}
