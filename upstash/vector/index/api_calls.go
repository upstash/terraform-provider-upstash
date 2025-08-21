package index

import (
	"github.com/upstash/terraform-provider-upstash/v2/upstash/client"
)

func CreateIndex(c *client.UpstashClient, body CreateIndexRequest) (index Index, err error) {

	resp, err := c.SendPostRequest("/v2/vector/index", body, "Create Vector Index", false)

	if err != nil {
		return index, err
	}

	err = resp.ToJSON(&index)
	return index, err

}

func GetIndex(c *client.UpstashClient, indexId string) (index Index, err error) {

	resp, err := c.SendGetRequest("/v2/vector/index/"+indexId, "Get Vector Index", false)

	if err != nil {
		return index, err
	}

	err = resp.ToJSON(&index)
	return index, err

}

func SetIndexPlan(c *client.UpstashClient, indexId string, plan SetPlanRequest) (err error) {
	_, err = c.SendPostRequest("/v2/vector/index/"+indexId+"/setplan", plan, "Set Plan for Vector Index", false)

	return err
}

func RenameIndex(c *client.UpstashClient, indexId string, name RenameIndexRequest) (err error) {
	_, err = c.SendPostRequest("/v2/vector/index/"+indexId+"/rename", name, "Rename Vector Index", false)

	return err
}

func DeleteIndex(c *client.UpstashClient, indexId string) (err error) {
	return c.SendDeleteRequest("/v2/vector/index/"+indexId, nil, "Delete Vector Index", false)
}
