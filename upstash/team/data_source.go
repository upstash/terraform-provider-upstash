package team

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceTeam() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceRead,
		Schema: map[string]*schema.Schema{
			"team_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Cluster ID for requested cluster",
			},
			"team_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the cluster",
			},
			"copy_cc": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether Credit card info is copied or not",
			},
		},
	}

}
