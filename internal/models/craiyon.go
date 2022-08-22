package models

import "fmt"

type CraiyonRequest struct {
	Prompt string `json:"prompt"`
}

func (c *CraiyonRequest) JSON() string {
	return fmt.Sprintf(`{"prompt":"%s"}`, c.Prompt)
}
