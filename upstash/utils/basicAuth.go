package utils

import "encoding/base64"

func BasicAuth(user string, password string) string {
	token := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(token))
}
