package client

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

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

func (c *UpstashClient) getQstashToken() (error, string) {
	type token struct {
		Token string `json:"token"`
	}
	resp, err := req.Get(
		UPSTASH_API_ENDPOINT+"/v2/qstash/user",
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)},
	)
	if err != nil {
		return err, ""
	}
	var qstashToken token
	err = resp.ToJSON(&qstashToken)

	return err, qstashToken.Token
}

func (c *UpstashClient) SendDeleteRequest(endpointExtensionOrQstashEndpoint string, body interface{}, errMessage string) (err error) {
	endpoint := UPSTASH_API_ENDPOINT + endpointExtensionOrQstashEndpoint
	if strings.Contains(endpointExtensionOrQstashEndpoint, "qstash") {
		err, BEARER_TOKEN := c.getQstashToken()
		if err != nil {
			return err
		}
		endpoint = endpointExtensionOrQstashEndpoint
		_, err = req.Delete(
			endpoint,
			req.Header{"Accept": "application/json"},
			req.Header{"Authorization": "Bearer " + BEARER_TOKEN},
			req.BodyJSON(body),
		)
		return err
	}
	resp, err := req.Delete(
		endpoint,
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

func (c *UpstashClient) SendGetRequest(endpointExtensionOrQstashEndpoint string, errMessage string) (response *req.Resp, err error) {
	endpoint := UPSTASH_API_ENDPOINT + endpointExtensionOrQstashEndpoint
	if strings.Contains(endpointExtensionOrQstashEndpoint, "qstash") {
		err, BEARER_TOKEN := c.getQstashToken()
		if err != nil {
			return response, err
		}
		endpoint = endpointExtensionOrQstashEndpoint
		return req.Get(
			endpoint,
			req.Header{"Accept": "application/json"},
			req.Header{"Authorization": "Bearer " + BEARER_TOKEN},
		)
	}
	resp, err := req.Get(
		endpoint,
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

func (c *UpstashClient) SendPostRequest(endpointExtensionOrQstashEndpoint string, body interface{}, errMessage string, headers ...req.Header) (response *req.Resp, err error) {

	endpoint := UPSTASH_API_ENDPOINT + endpointExtensionOrQstashEndpoint
	if strings.Contains(endpointExtensionOrQstashEndpoint, "qstash") {
		err, BEARER_TOKEN := c.getQstashToken()
		if err != nil {
			return response, err
		}
		endpoint = endpointExtensionOrQstashEndpoint
		return req.Post(
			endpoint,
			req.Header{"Accept": "application/json"},
			req.Header{"Authorization": "Bearer " + BEARER_TOKEN},
			req.Header{"Upstash-Cron": "* * * * *"},
			// headers,
			req.BodyJSON(body),
		)
	}
	resp, err := req.Post(
		endpoint,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)},
		req.BodyJSON(body),
		headers,
	)

	if err != nil {
		return nil, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return nil, errors.New(errMessage + " failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return resp, err
}
