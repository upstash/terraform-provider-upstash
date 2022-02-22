package client

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/imroc/req"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
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

func (c *UpstashClient) SendDeleteRequest(endpointExtension string, body interface{}, errMessage string) (err error) {
	resp, err := req.Delete(
		UPSTASH_API_ENDPOINT+endpointExtension,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)},
		req.BodyJSON(body),
	)
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New(errMessage + " failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return err
}

func (c *UpstashClient) SendGetRequest(endpointExtension string, errMessage string) (response *req.Resp, err error) {
	resp, err := req.Get(
		UPSTASH_API_ENDPOINT+endpointExtension,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)},
	)
	if err != nil {
		return resp, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return resp, errors.New(errMessage + " failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return resp, err
}

func (c *UpstashClient) SendPostRequest(endpointExtension string, body interface{}, errMessage string) (response *req.Resp, err error) {

	resp, err := req.Post(
		UPSTASH_API_ENDPOINT+endpointExtension,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)},
		req.BodyJSON(body),
	)

	if err != nil {
		return nil, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return nil, errors.New(errMessage + " failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return resp, err
}
