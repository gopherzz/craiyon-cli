package convert

import (
	"bytes"
	"encoding/base64"
	"image"
	"io/ioutil"
	"strings"

	"github.com/gopherzz/craiyon-api/internal/config"
	"github.com/gopherzz/craiyon-api/internal/models"
	"github.com/nfnt/resize"
)

func ImagesResponseToImages(resp models.ImagesResponse, cfg config.Config) (models.Images, error) {
	var imagesBytes models.Images
	for _, imageBytes := range resp.Images {
		png, err := base64ToPng(imageBytes)
		if err != nil {
			return imagesBytes, err
		}
		// convert bytes to image
		img, _, err := image.Decode(bytes.NewReader(png))
		if err != nil {
			return imagesBytes, err
		}
		img = resize.Resize(512, 512, img, resize.Lanczos3)
		imagesBytes = append(imagesBytes, img)
	}
	return imagesBytes, nil
}

func base64ToPng(base64String string) ([]byte, error) {
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64String))
	return ioutil.ReadAll(decoder)
}
