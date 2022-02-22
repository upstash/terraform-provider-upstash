package integrationtesting

import "os"

type UpstashCredentials struct {
	Email  string
	Apikey string
}

func GetCredentials() UpstashCredentials {
	return UpstashCredentials{
		Email:  os.Getenv("UPSTASH_EMAIL"),
		Apikey: os.Getenv("UPSTASH_API_KEY"),
	}
}
