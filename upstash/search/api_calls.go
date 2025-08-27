package search

import (
	"github.com/upstash/terraform-provider-upstash/v2/upstash/client"
)

func CreateSearch(c *client.UpstashClient, body CreateSearchRequest) (search Search, err error) {
	resp, err := c.SendPostRequest("/v2/search", body, "Create Search", false)
	if err != nil {
		return search, err
	}

	err = resp.ToJSON(&search)
	return search, err
}

func GetSearch(c *client.UpstashClient, searchId string) (search Search, err error) {
	resp, err := c.SendGetRequest("/v2/search/"+searchId, "Get Search", false)
	if err != nil {
		return search, err
	}

	err = resp.ToJSON(&search)
	return search, err
}

func SetSearchPlan(c *client.UpstashClient, searchId string, plan SetPlanRequest) (err error) {
	_, err = c.SendPostRequest("/v2/search/"+searchId+"/setplan", plan, "Set Plan for Search", false)
	return err
}

func RenameSearch(c *client.UpstashClient, searchId string, name RenameSearchRequest) (err error) {
	_, err = c.SendPostRequest("/v2/search/"+searchId+"/rename", name, "Rename Search", false)
	return err
}

func DeleteSearch(c *client.UpstashClient, searchId string) (err error) {
	return c.SendDeleteRequest("/v2/search/"+searchId, nil, "Delete Search", false)
}
