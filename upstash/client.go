package upstash

import (
	"encoding/base64"
	"errors"
	"github.com/imroc/req"
	"net/http"
	"strconv"
)

const UPSTASH_API_ENDPOINT = "https://api.upstash.com"

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

func (c *UpstashClient) CreateDatabase(body CreateDatabaseRequest) (database Database, err error) {
	resp, err := req.Post(UPSTASH_API_ENDPOINT+"/v1/database",
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": basicAuth(c.Email, c.Apikey)},
		req.BodyJSON(body))
	if err != nil {
		return database, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return database, errors.New("Create database failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&database)
	return database, err
}

func (c *UpstashClient) GetDatabase(databaseId string) (database Database, err error) {
	resp, err := req.Get(UPSTASH_API_ENDPOINT+"/v1/database/"+databaseId,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": basicAuth(c.Email, c.Apikey)})
	if err != nil {
		return database, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return database, errors.New("Get database failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&database)
	return database, err
}

func (c *UpstashClient) DeleteDatabase(databaseId string) (err error) {
	resp, err := req.Delete(UPSTASH_API_ENDPOINT+"/v1/database/"+databaseId,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": basicAuth(c.Email, c.Apikey)})
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Get database failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return err
}

func (c *UpstashClient) EnableTLS(databaseId string) (err error) {
	resp, err := req.Post(UPSTASH_API_ENDPOINT+"/v1/tls/"+databaseId,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": basicAuth(c.Email, c.Apikey)})
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Enable TLS failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return err
}

func (c *UpstashClient) EnableMultiZone(databaseId string, enabled bool) (err error) {
	resp, err := req.Post(UPSTASH_API_ENDPOINT+"/v1/multizone/"+databaseId,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": basicAuth(c.Email, c.Apikey)}, req.BodyJSON(req.Param{"enabled": enabled}))
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Enable TLS failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return err
}

func basicAuth(user string, password string) string {
	token := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(token))
}
