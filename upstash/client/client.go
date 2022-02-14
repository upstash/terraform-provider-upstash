package client

const UPSTASH_API_ENDPOINT = "https://api.upstash.com"

// const UPSTASH_V2_API_ENDPOINT = "https://api-dev.upstash.io"

type UpstashClient struct {
	Email  string
	Apikey string
}

func NewUpstashClient(email string, apikey string) *UpstashClient {
	return &UpstashClient{
		Email:  email,
		Apikey: apikey,
	}
}
