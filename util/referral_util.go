package util

import (
	"encoding/base64"
)

func GenerateReferral(name string, email string) string {
	data := name + email
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	return encodedData
}
