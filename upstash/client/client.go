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

func (c *UpstashClient) GetQstashEndpoint() string {
	return "https://qstash.upstash.io/v1"
}

func (c *UpstashClient) GetQstashToken() (error, string) {
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

func (c *UpstashClient) SendDeleteRequest(endpointExtensionOrQstashEndpoint string, body interface{}, errMessage string, qstashFlag bool) (err error) {
	endpoint := UPSTASH_API_ENDPOINT + endpointExtensionOrQstashEndpoint
	authHeader := req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)}

	if qstashFlag {
		endpoint = endpointExtensionOrQstashEndpoint
		err, BEARER_TOKEN := c.GetQstashToken()
		if err != nil {
			return err
		}
		authHeader = req.Header{"Authorization": "Bearer " + BEARER_TOKEN}
	}

	resp, err := req.Delete(
		endpoint,
		req.Header{"Accept": "application/json"},
		authHeader,
		req.BodyJSON(body),
	)
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted && resp.Response().StatusCode != http.StatusCreated {
		return errors.New(errMessage + " failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return err
}

func (c *UpstashClient) SendGetRequest(endpointExtensionOrQstashEndpoint string, errMessage string, qstashFlag bool) (response *req.Resp, err error) {
	endpoint := UPSTASH_API_ENDPOINT + endpointExtensionOrQstashEndpoint
	authHeader := req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)}

	if qstashFlag {
		endpoint = endpointExtensionOrQstashEndpoint
		err, BEARER_TOKEN := c.GetQstashToken()
		if err != nil {
			return response, err
		}
		authHeader = req.Header{"Authorization": "Bearer " + BEARER_TOKEN}
	}

	resp, err := req.Get(
		endpoint,
		req.Header{"Accept": "application/json"},
		authHeader,
	)
	if err != nil {
		return resp, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted && resp.Response().StatusCode != http.StatusCreated {
		return resp, errors.New(errMessage + " failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return resp, err
}

func (c *UpstashClient) SendPostRequest(endpointExtensionOrQstashEndpoint string, body interface{}, errMessage string, qstashFlag bool) (response *req.Resp, err error) {
	endpoint := UPSTASH_API_ENDPOINT + endpointExtensionOrQstashEndpoint
	authHeader := req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)}

	if qstashFlag {
		endpoint = endpointExtensionOrQstashEndpoint
		err, BEARER_TOKEN := c.GetQstashToken()
		if err != nil {
			return response, err
		}
		authHeader = req.Header{"Authorization": "Bearer " + BEARER_TOKEN}
	}

	resp, err := req.Post(
		endpoint,
		req.Header{"Accept": "application/json"},
		authHeader,
		req.BodyJSON(body),
	)

	if err != nil {
		return nil, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted && resp.Response().StatusCode != http.StatusCreated {
		return nil, errors.New(errMessage + " failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return resp, err
}

func (c *UpstashClient) SendPutRequest(endpointExtensionOrQstashEndpoint string, body interface{}, errMessage string, qstashFlag bool) (response *req.Resp, err error) {
	endpoint := UPSTASH_API_ENDPOINT + endpointExtensionOrQstashEndpoint
	authHeader := req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)}

	if qstashFlag {
		endpoint = endpointExtensionOrQstashEndpoint
		err, BEARER_TOKEN := c.GetQstashToken()
		if err != nil {
			return response, err
		}
		authHeader = req.Header{"Authorization": "Bearer " + BEARER_TOKEN}
	}

	resp, err := req.Put(
		endpoint,
		req.Header{"Accept": "application/json"},
		authHeader,
		req.BodyJSON(body),
	)

	if err != nil {
		return nil, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted && resp.Response().StatusCode != http.StatusCreated {
		return nil, errors.New(errMessage + " failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return resp, err
}

func (c *UpstashClient) SendPatchRequest(endpointExtensionOrQstashEndpoint string, body interface{}, errMessage string, qstashFlag bool) (response *req.Resp, err error) {
	endpoint := UPSTASH_API_ENDPOINT + endpointExtensionOrQstashEndpoint
	authHeader := req.Header{"Authorization": utils.BasicAuth(c.Email, c.Apikey)}

	if qstashFlag {
		endpoint = endpointExtensionOrQstashEndpoint
		err, BEARER_TOKEN := c.GetQstashToken()
		if err != nil {
			return response, err
		}
		authHeader = req.Header{"Authorization": "Bearer " + BEARER_TOKEN}
	}

	resp, err := req.Patch(
		endpoint,
		req.Header{"Accept": "application/json"},
		authHeader,
		req.BodyJSON(body),
	)

	if err != nil {
		return nil, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted && resp.Response().StatusCode != http.StatusCreated {
		return nil, errors.New(errMessage + " failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return resp, err
}
