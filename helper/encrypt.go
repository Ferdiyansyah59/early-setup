package helper

import "encoding/base64"

func GenerateEncrypt(word string) string {
	return base64.StdEncoding.EncodeToString([]byte(word))
}