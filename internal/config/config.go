package config

import (
	"github.com/gopherzz/craiyon-api/internal/models"
)

type Config struct {
	Request models.CraiyonRequest `json:"prompt"`
	Width   int                   `json:"width"`
	Height  int                   `json:"height"`
	Quality int                   `json:"quality"`
}

var Default = Config{
	Request: models.CraiyonRequest{},
	Width:   512,
	Height:  512,
	Quality: 100,
}
