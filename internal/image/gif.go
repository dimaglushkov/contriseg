package image

import (
	"image"
	"image/gif"
	"os"
)

func GenerateGIF(frames []*image.Paletted, filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0633)
	if err != nil {
		return err
	}
	defer f.Close()
	err = gif.EncodeAll(f, &gif.GIF{
		Image: frames,
		Delay: make([]int, len(frames)),
	})
	return err
}
