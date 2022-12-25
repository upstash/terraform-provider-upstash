package database

import (
	"github.com/upstash/terraform-provider-upstash/upstash/client"
)

// not needed if changes

func CreateDatabase(c *client.UpstashClient, body CreateDatabaseRequest) (database Database, err error) {

	resp, err := c.SendPostRequest("/v2/redis/database", body, "Create Redis Database", false)

	if err != nil {
		return database, err
	}

	err = resp.ToJSON(&database)
	return database, err

}

func GetDatabase(c *client.UpstashClient, databaseId string) (database Database, err error) {

	resp, err := c.SendGetRequest("/v2/redis/database/"+databaseId, "Get Redis Database", false)

	if err != nil {
		return database, err
	}

	err = resp.ToJSON(&database)
	return database, err

}

func EnableTLS(c *client.UpstashClient, databaseId string) (err error) {

	_, err = c.SendPostRequest("/v2/redis/enable-tls/"+databaseId, nil, "Enable Tls for Redis Database", false)

	return err

}

func EnableMultiZone(c *client.UpstashClient, databaseId string, enabled bool) (err error) {

	_, err = c.SendPostRequest("/v2/redis/enable-multizone/"+databaseId, nil, "Enable Multizone for Redis Database", false)

	return err

}

func ConfigureEviction(c *client.UpstashClient, databaseId string, enabled bool) (err error) {

	body := ConfigureEvictionRequest{
		DatabaseId: databaseId,
		Eviction:   enabled,
	}
	_, err = c.SendPatchRequest("/v2/redis/eviction", body, "Configure Eviction Redis Database", false)

	return err

}

func ConfigureAutoUpgrade(c *client.UpstashClient, databaseId string, enabled bool) (err error) {

	body := ConfigureAutoUpgradeRequest{
		DatabaseId:  databaseId,
		AutoUpgrade: enabled,
	}
	_, err = c.SendPatchRequest("/v2/redis/autoupgrade", body, "Configure Auto Upgrade for Redis Database", false)

	return err

}

func DeleteDatabase(c *client.UpstashClient, databaseId string) (err error) {

	return c.SendDeleteRequest("/v2/redis/database/"+databaseId, nil, "Delete Redis Database", false)

}
