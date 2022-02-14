package database

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/imroc/req"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

const api_endpoint = client.UPSTASH_API_ENDPOINT

func CreateDatabase(c *client.UpstashClient, body CreateDatabaseRequest) (database Database, err error) {
	resp, err := req.Post(api_endpoint+"/v2/redis/database",
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)},
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

func GetDatabase(c *client.UpstashClient, databaseId string) (database Database, err error) {
	resp, err := req.Get(api_endpoint+"/v2/redis/database/"+databaseId,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)})
	if err != nil {
		return database, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return database, errors.New("Get database failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&database)
	return database, err
}

func EnableTLS(c *client.UpstashClient, databaseId string) (err error) {
	resp, err := req.Post(api_endpoint+"/v2/redis/enable-tls/"+databaseId,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)})
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Enable TLS failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return err
}

func EnableMultiZone(c *client.UpstashClient, databaseId string, enabled bool) (err error) {
	resp, err := req.Post(api_endpoint+"/v2/redis/enable-multizone/"+databaseId,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)}, req.BodyJSON(req.Param{"enabled": enabled}))
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Enable Multizone failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return err
}

func DeleteDatabase(c *client.UpstashClient, databaseId string) (err error) {
	resp, err := req.Delete(api_endpoint+"/v2/redis/database/"+databaseId,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)})
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Get database failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return err
}
