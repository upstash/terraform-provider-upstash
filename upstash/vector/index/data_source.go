package index

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataResourceIndex() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceIndexRead,

		Schema: map[string]*schema.Schema{
			"customer_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique ID associated to the owner of this index.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Index ID for created index.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the index.",
			},
			"similarity_function": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Associated distance metric to calculate the similarity.",
			},
			"dimension_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Size of the vector array.",
			},
			"endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Associated endpoint of your index.",
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "REST token to send request to the related index.",
			},
			"read_only_token": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Readonly REST token to send request to the related index. You can't perform update operation with this token.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Associated plan of the index. Either `free`, `paid`, `fixed` or `pro`.",
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The region where your index is deployed.",
			},
			"max_vector_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum that your index can contain.",
			},
			"max_daily_updates": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum update operations you can perform in a day. Only upsert operations are included in update count.",
			},
			"max_daily_queries": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum query operations you can perform in a day. Only query operations are included in query count.",
			},
			"max_monthly_bandwidth": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum amount of monthly bandwidth for the index. Unit is bytes. `-1` if the limit is unlimited.",
			},
			"max_writes_per_second": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum write operations you can perform per second. Only upsert operations are included in write count.",
			},
			"max_query_per_second": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum query operations you can perform per second. Only query operations are included in query count.",
			},
			"max_reads_per_request": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum vectors in a read operation. Query and fetch operations are included in read operations.",
			},
			"max_writes_per_request": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of maximum vectors in a write operation. Only upsert operations are included in write operations.",
			},
			"max_total_metadata_size": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The amount of maximum size for the total metadata sizes in your index.",
			},
			"reserved_price": &schema.Schema{
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Monthly pricing of your index. Only available for fixed and pro plans.",
			},
			"creation_time": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The creation time of the vector index in UTC as unix timestamp.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
