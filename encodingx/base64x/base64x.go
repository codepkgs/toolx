package base64x

import (
	"encoding/base64"
)

func EncodeString(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func DecodeStr(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func EncodeStringWithURLEncoding(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

func DecodeStringWithURLEncoding(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}
