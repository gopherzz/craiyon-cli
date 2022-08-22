package generate

import (
	"encoding/json"
	"log"

	"github.com/gopherzz/craiyon-api/internal/config"
	"github.com/gopherzz/craiyon-api/internal/convert"
	"github.com/gopherzz/craiyon-api/internal/models"
	"github.com/gopherzz/craiyon-api/pkg/craiyon"
)

func GenerateImages(cfg config.Config) (models.Images, error) {
	imagesJson, err := craiyon.GenerateImagesJson(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	var images models.ImagesResponse
	err = json.NewDecoder(imagesJson).Decode(&images)
	if err != nil {
		return nil, err
	}

	return convert.ImagesResponseToImages(images, cfg)
}
