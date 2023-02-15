package image

import (
	"image"
	"image/gif"
	"os"
	"path/filepath"
)

func GenerateGIF(frames []*image.Paletted, filename string) error {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
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
