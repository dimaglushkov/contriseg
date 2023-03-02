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

var ColorScheme = [...]color.RGBA{
	GithubNoContribColor:      {24, 28, 36, 1},
	GithubFirstQuadrantColor:  {14, 68, 41, 1},
	GithubSecondQuadrantColor: {0, 109, 50, 1},
	GithubThirdQuadrantColor:  {38, 166, 65, 1},
	GithubFourthQuadrantColor: {57, 211, 83, 1},
	BackgroundColor:           {13, 17, 23, 1},
	DarkBlueColor:             {16, 33, 97, 1},
	BlueColor:                 {10, 30, 150, 1},
}
