package connector

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceConnector() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceReadForDataSource,
		Schema: map[string]*schema.Schema{
			"connector_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Connector ID for created connector",
			},
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Cluster ID for cluster that the connector is tied to",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the connector",
			},
			"properties": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Properties that the connector will have",
			},
			"creation_time": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation of the connector",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of the connector",
			},
			"state_error_message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State Error Message of the connector",
			},
			"connector_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State Error Message of the connector",
			},
			"connector_class": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Connector class of the connector",
			},
			"encoded_username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Encoded username for the connector",
			},
			"properties_encrypted": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Encrypted properties for the connector",
			},
			"user_password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "User password for the connector",
			},
			"ttl": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "TTL for the connector",
			},
			"topics": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Topics for the connector",
			},
			"tasks": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
				Description: "Tasks of the connector",
			},
		},
	}
}
