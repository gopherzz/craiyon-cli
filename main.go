package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"

	"github.com/gopherzz/craiyon-api/internal/config"
	"github.com/gopherzz/craiyon-api/internal/generate"
	"github.com/gopherzz/craiyon-api/internal/models"
)

func saveImages(images models.Images, cfg config.Config) error {
	err := os.Mkdir(cfg.Request.Prompt, 0755)
	if err != nil {
		return err
	}

	for i, image := range images {
		f, err := os.OpenFile(fmt.Sprintf("%s/%s-%d.png", cfg.Request.Prompt, cfg.Request.Prompt, i), os.O_WRONLY|os.O_CREATE, 0666)
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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a prompt")
	}

	prompt := os.Args[1]

	cfg := config.Default
	cfg.Request.Prompt = prompt

	images, err := generate.GenerateImages(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = saveImages(images, cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
}
