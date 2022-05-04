package helpers

import "encoding/base64"

func ConvertBase64ToString(s string) string {

	data, err := base64.StdEncoding.DecodeString(s)

	CheckError(err)

	return string(data)
}
