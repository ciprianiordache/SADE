package utility

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateLink() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	link := base64.URLEncoding.EncodeToString(token)
	return link, nil
}
