package models

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

type ImagesResponse struct {
	Images  []string `json:"images"`
	Version string   `json:"version"`
}

type Images []image.Image

func (images Images) saveWithName(prompt string) error {
	err := os.Mkdir(prompt, 0755)
	if err != nil {
		return err
	}

	for i, image := range images {
		f, err := os.OpenFile(fmt.Sprintf("%s/%s-%d.png", prompt, prompt, i), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer f.Close()

		err = jpeg.Encode(f, image, &jpeg.Options{Quality: 100})
		if err != nil {
			return err
		}
	}

	return nil
}
