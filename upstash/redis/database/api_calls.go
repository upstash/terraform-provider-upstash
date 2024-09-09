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

func UpdateReadRegions(c *client.UpstashClient, databaseId string, readRegions UpdateReadRegionsRequest) (err error) {

	_, err = c.SendPostRequest("/v2/redis/update-regions/"+databaseId, readRegions, "Update Regions for Redis Database", false)

	return err

}

func EnableTLS(c *client.UpstashClient, databaseId string) (err error) {

	_, err = c.SendPostRequest("/v2/redis/enable-tls/"+databaseId, nil, "Enable Tls for Redis Database", false)

	return err

}

func ConfigureEviction(c *client.UpstashClient, databaseId string, enabled bool) (err error) {
	path := "/v2/redis/"
	if enabled {
		path += "enable-eviction/"
	} else {
		path += "disable-eviction/"
	}
	path += databaseId

	_, err = c.SendPostRequest(path, nil, "Configure Eviction Redis Database", false)

	return err
}

func ConfigureAutoUpgrade(c *client.UpstashClient, databaseId string, enabled bool) (err error) {

	path := "/v2/redis/"
	if enabled {
		path += "enable-autoupgrade/"
	} else {
		path += "disable-autoupgrade/"
	}
	path += databaseId
	_, err = c.SendPostRequest(path, nil, "Configure Auto Upgrade for Redis Database", false)

	return err

}

func DeleteDatabase(c *client.UpstashClient, databaseId string) (err error) {

	return c.SendDeleteRequest("/v2/redis/database/"+databaseId, nil, "Delete Redis Database", false)

}
