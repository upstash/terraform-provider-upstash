package database

import "github.com/upstash/terraform-provider-upstash/upstash/client"

type upstashClient client.UpstashClient

type Database struct {
	DatabaseId   string `json:"database_id"`
	DatabaseName string `json:"database_name"`
	Region       string `json:"region"`
	Replicas     int    `json:"replicas"`
	Port         int    `json:"port"`
	CreationTime int64  `json:"creation_time"`
	Password     string `json:"password,omitempty"`
	User         string `json:"customer_id"`
	Endpoint     string `json:"endpoint"`
	Tls          bool   `json:"tls"`
	Consistent   bool   `json:"consistent"`
	MultiZone    bool   `json:"multizone"`
	RestToken    string `json:"rest_token,omitempty"`
}

type CreateDatabaseRequest struct {
	Region       string `json:"region"`
	DatabaseName string `json:"database_name"`
	Tls          bool   `json:"tls"`
	Consistent   bool   `json:"consistent"`
	MultiZone    bool   `json:"multizone"`
}
