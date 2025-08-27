package search

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataResourceSearch() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceSearchRead,

		Schema: map[string]*schema.Schema{
			"customer_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique ID associated to the owner of this search.",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Search ID for created search.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the search.",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Associated endpoint of your search.",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "REST token to send request to the related search.",
			},
			"read_only_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Readonly REST token to send request to the related search. You can't perform update operation with this token.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Associated plan of the search. Either `free`, `paid`, `fixed` or `pro`.",
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The region where your search is deployed.",
			},
			"max_vector_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum that your search can contain.",
			},
			"max_daily_updates": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum update operations you can perform in a day. Only upsert operations are included in update count.",
			},
			"max_daily_queries": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum query operations you can perform in a day. Only query operations are included in query count.",
			},
			"max_monthly_bandwidth": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum amount of monthly bandwidth for the search. Unit is bytes. `-1` if the limit is unlimited.",
			},
			"max_writes_per_second": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum write operations you can perform per second. Only upsert operations are included in write count.",
			},
			"max_query_per_second": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum query operations you can perform per second. Only query operations are included in query count.",
			},
			"max_reads_per_request": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum vectors in a read operation. Query and fetch operations are included in read operations.",
			},
			"max_writes_per_request": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum vectors in a write operation. Only upsert operations are included in write operations.",
			},
			"max_total_metadata_size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The amount of maximum size for the total metadata sizes in your search.",
			},
			"reserved_price": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Monthly pricing of your search. Only available for fixed and pro plans.",
			},
			"creation_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The creation time of the vector search in UTC as unix timestamp.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
