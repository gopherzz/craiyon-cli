package craiyon

import (
	"io"
	"net/http"
	"strings"

	"github.com/gopherzz/craiyon-api/internal/config"
)

func GenerateImagesJson(cfg config.Config) (io.Reader, error) {
	client := http.Client{}
	resp, err := client.Post("https://backend.craiyon.com/generate", "application/json", strings.NewReader(cfg.Request.JSON()))
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
