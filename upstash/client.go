package upstash

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"

	"github.com/imroc/req"
	"github.com/upstash/terraform-provider-upstash/upstash/kafka/cluster"
	"github.com/upstash/terraform-provider-upstash/upstash/redis/database"
)

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

func (c *UpstashClient) CreateDatabase(body database.CreateDatabaseRequest) (database database.Database, err error) {
	resp, err := req.Post(UPSTASH_API_ENDPOINT+"/v2/redis/database",
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

func (c *UpstashClient) GetDatabase(databaseId string) (database database.Database, err error) {
	resp, err := req.Get(UPSTASH_API_ENDPOINT+"/v2/redis/database/"+databaseId,
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
	resp, err := req.Delete(UPSTASH_API_ENDPOINT+"/v2/redis/database/"+databaseId,
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
	resp, err := req.Post(UPSTASH_API_ENDPOINT+"/v2/redis/enable-tls/"+databaseId,
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
	resp, err := req.Post(UPSTASH_API_ENDPOINT+"/v2/redis/enable-multizone/"+databaseId,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": basicAuth(c.Email, c.Apikey)}, req.BodyJSON(req.Param{"enabled": enabled}))
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Enable Multizone failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return err
}

func basicAuth(user string, password string) string {
	token := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(token))
}

// Clusters

func (c *UpstashClient) DeleteCluster(clusterId string) (err error) {
	resp, err := req.Delete(UPSTASH_API_ENDPOINT+"/v2/kafka/cluster/"+clusterId,
		req.Header{
			"Accept":        "application/json",
			"Authorization": basicAuth(c.Email, c.Apikey),
		},
	)
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Delete cluster failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())

	}
	return err
}

func (c *UpstashClient) GetCluster(clusterId string) (cluster cluster.Cluster, err error) {
	resp, err := req.Get(UPSTASH_API_ENDPOINT+"/v2/kafka/cluster/"+clusterId,
		req.Header{
			"Accept":        "application/json",
			"Authorization": basicAuth(c.Email, c.Apikey),
		},
	)
	if err != nil {
		return cluster, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return cluster, errors.New("Get cluster failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&cluster)
	return cluster, err

}

func (c *UpstashClient) CreateCluster(body cluster.CreateClusterRequest) (cluster cluster.Cluster, err error) {
	resp, err := req.Post(UPSTASH_API_ENDPOINT+"/v2/kafka/cluster",
		req.Header{
			"Accept":        "application/json",
			"Authorization": basicAuth(c.Email, c.Apikey),
		},
		req.BodyJSON(body),
	)
	if err != nil {
		return cluster, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return cluster, errors.New("Create cluster failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&cluster)
	return cluster, err

}

// how to send as body parameter?
func (c *UpstashClient) RenameCluster(clusterId string, newName string) (err error) {

	header := req.Header{
		"Accept":        "application/json",
		"Authorization": basicAuth(c.Email, c.Apikey),
	}

	param := req.Param{
		"name": newName,
	}

	resp, err := req.Post(UPSTASH_API_ENDPOINT+"/v2/kafka/rename-cluster/"+clusterId, req.BodyJSON(param), header)

	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Renaming cluster failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return nil
}
